// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package main

import (
	"strings"
	"testing"
)

func process(t *testing.T, schema string, depth int) (string, []string) {
	t.Helper()

	rewritten, notes, err := processSchema([]byte(schema), depth)
	if err != nil {
		t.Fatalf("processSchema: %s", err)
	}
	return string(rewritten), notes
}

func TestNoChanges(t *testing.T) {
	schema := `{
  "typeName": "AWS::Test::Test",
  "definitions": {
    "A": {
      "type": "object",
      "properties": {
        "B": {
          "$ref": "#/definitions/B"
        }
      }
    },
    "B": {
      "type": "string"
    }
  }
}
`
	rewritten, notes := process(t, schema, 3)

	if rewritten != "" {
		t.Errorf("expected no rewrite, got:\n%s", rewritten)
	}
	if len(notes) != 0 {
		t.Errorf("expected no notes, got %v", notes)
	}
}

func TestAliasInlining(t *testing.T) {
	schema := `{
  "definitions": {
    "Alias": {
      "$ref": "#/definitions/Target"
    },
    "TransitiveAlias": {
      "$ref": "#/definitions/Alias"
    },
    "Target": {
      "type": "string"
    },
    "User": {
      "type": "array",
      "items": {
        "$ref": "#/definitions/TransitiveAlias"
      }
    }
  }
}
`
	rewritten, notes := process(t, schema, 3)

	want := `{
  "definitions": {
    "Alias": {
      "$ref": "#/definitions/Target"
    },
    "TransitiveAlias": {
      "$ref": "#/definitions/Alias"
    },
    "Target": {
      "type": "string"
    },
    "User": {
      "type": "array",
      "items": {
        "$ref": "#/definitions/Target"
      }
    }
  }
}
`
	if rewritten != want {
		t.Errorf("unexpected rewrite:\n%s", rewritten)
	}
	wantNotes := []string{"alias TransitiveAlias inlined to Target"}
	if len(notes) != len(wantNotes) || notes[0] != wantNotes[0] {
		t.Errorf("unexpected notes: %v", notes)
	}
}

// TestUnroll exercises the WAFv2-like shape: a counting member (Statement)
// with a required-recursion branch (AndStatement, unconstructible at the
// last level) and an optional-recursion branch (RateStatement, kept at the
// last level minus its optional recursive property).
func TestUnroll(t *testing.T) {
	schema := `{
  "definitions": {
    "Statement": {
      "type": "object",
      "properties": {
        "AndStatement": {
          "$ref": "#/definitions/AndStatement"
        },
        "RateStatement": {
          "$ref": "#/definitions/RateStatement"
        },
        "Match": {
          "type": "string"
        }
      }
    },
    "AndStatement": {
      "type": "object",
      "properties": {
        "Statements": {
          "type": "array",
          "items": {
            "$ref": "#/definitions/Statement"
          }
        }
      },
      "required": [
        "Statements"
      ]
    },
    "RateStatement": {
      "type": "object",
      "properties": {
        "ScopeDownStatement": {
          "$ref": "#/definitions/Statement"
        }
      }
    }
  }
}
`
	rewritten, notes := process(t, schema, 2)

	wantNote := "cycle {AndStatement, RateStatement, Statement} unrolled to depth 2 (counting refs into Statement; 2 cloned definitions added)"
	if len(notes) != 1 || notes[0] != wantNote {
		t.Fatalf("unexpected notes: %v", notes)
	}

	// Only refs into the counting member (Statement) increment the level, so
	// level-1 AndStatement/RateStatement point at StatementLevel2. A level-2
	// AndStatement would need level-3 required Statements, so it is
	// unconstructible and dropped from StatementLevel2 (and never emitted);
	// level-2 RateStatement survives minus its optional ScopeDownStatement.
	// Both pruned property names are recorded for the code generator.
	want := `{
  "definitions": {
    "Statement": {
      "type": "object",
      "properties": {
        "AndStatement": {
          "$ref": "#/definitions/AndStatement"
        },
        "RateStatement": {
          "$ref": "#/definitions/RateStatement"
        },
        "Match": {
          "type": "string"
        }
      }
    },
    "AndStatement": {
      "type": "object",
      "properties": {
        "Statements": {
          "type": "array",
          "items": {
            "$ref": "#/definitions/StatementLevel2"
          }
        }
      },
      "required": [
        "Statements"
      ]
    },
    "RateStatement": {
      "type": "object",
      "properties": {
        "ScopeDownStatement": {
          "$ref": "#/definitions/StatementLevel2"
        }
      }
    },
    "RateStatementLevel2": {
      "type": "object",
      "properties": {}
    },
    "StatementLevel2": {
      "type": "object",
      "properties": {
        "RateStatement": {
          "$ref": "#/definitions/RateStatementLevel2"
        },
        "Match": {
          "type": "string"
        }
      }
    }
  },
  "x-derecursed": {
    "depth": 2,
    "prunedProperties": [
      "AndStatement",
      "ScopeDownStatement"
    ]
  }
}
`
	if rewritten != want {
		t.Errorf("unexpected rewrite:\n%s", rewritten)
	}

	// Idempotency: a second pass finds nothing to do.
	second, secondNotes, err := processSchema([]byte(rewritten), 2)
	if err != nil {
		t.Fatalf("second pass: %s", err)
	}
	if second != nil || len(secondNotes) != 0 {
		t.Errorf("second pass was not a no-op: %v\n%s", secondNotes, second)
	}
}

func TestUnconstructibleAtLevelOne(t *testing.T) {
	schema := `{
  "definitions": {
    "A": {
      "type": "object",
      "properties": {
        "A": {
          "$ref": "#/definitions/A"
        }
      },
      "required": [
        "A"
      ]
    }
  }
}
`
	_, _, err := processSchema([]byte(schema), 3)

	want := "error: A is unconstructible even at level 1; increase --depth"
	if err == nil || err.Error() != want {
		t.Errorf("expected %q, got %v", want, err)
	}
}

// TestAliasCycle also pins the reported member: with more than one alias cycle
// the error must not depend on map iteration order.
func TestAliasCycle(t *testing.T) {
	schema := `{
  "definitions": {
    "A": {
      "$ref": "#/definitions/B"
    },
    "B": {
      "$ref": "#/definitions/A"
    },
    "C": {
      "$ref": "#/definitions/D"
    },
    "D": {
      "$ref": "#/definitions/C"
    }
  }
}
`
	want := "error: alias cycle at A"
	for range 20 {
		_, _, err := processSchema([]byte(schema), 3)

		if err == nil || err.Error() != want {
			t.Fatalf("expected %q, got %v", want, err)
		}
	}
}

// TestKeyOrderPreserved verifies that untouched parts of the schema
// round-trip with their original key order and formatting.
func TestKeyOrderPreserved(t *testing.T) {
	schema := `{
  "zebra": {
    "b": 1,
    "a": [
      true,
      null,
      1.5
    ]
  },
  "alpha": "text with \"quotes\" and <angle> & ampersand",
  "definitions": {
    "Alias": {
      "$ref": "#/definitions/Target"
    },
    "Target": {
      "type": "string"
    },
    "User": {
      "$comment": "forces a rewrite so the file is re-serialized",
      "items": {
        "$ref": "#/definitions/Alias"
      }
    }
  }
}
`
	rewritten, _ := process(t, schema, 3)

	want := strings.Replace(schema, `"$ref": "#/definitions/Alias"`, `"$ref": "#/definitions/Target"`, 1)
	if rewritten != want {
		t.Errorf("round-trip changed unrelated content:\ngot:\n%s\nwant:\n%s", rewritten, want)
	}
}
