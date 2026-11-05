// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package codegen

import (
	"os"
	"path/filepath"
	"testing"
)

func TestUnrepresentableProperties(t *testing.T) {
	t.Parallel()

	dir := t.TempDir()
	write := func(name, body string) string {
		p := filepath.Join(dir, name)
		if err := os.WriteFile(p, []byte(body), 0o600); err != nil {
			t.Fatal(err)
		}
		return p
	}

	pristine := write("pristine.json", `{"typeName":"AWS::Test::Type","properties":{}}`)
	got, err := unrepresentableProperties(pristine)
	if err != nil {
		t.Fatalf("pristine schema: %v", err)
	}
	if got != nil {
		t.Errorf("pristine schema: want nil, got %v", got)
	}

	derecursed := write("derecursed.json", `{"typeName":"AWS::Test::Type","properties":{},"x-derecursed":{"depth":4,"prunedProperties":["Statement","ScopeDownStatement"]}}`)
	got, err = unrepresentableProperties(derecursed)
	if err != nil {
		t.Fatalf("de-recursed schema: %v", err)
	}
	if len(got) != 2 || got[0] != "Statement" || got[1] != "ScopeDownStatement" {
		t.Errorf("de-recursed schema: want [Statement ScopeDownStatement], got %v", got)
	}

	if _, err := unrepresentableProperties(write("bad.json", `{`)); err == nil {
		t.Error("malformed schema: want error, got nil")
	}
}
