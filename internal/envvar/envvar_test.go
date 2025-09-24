// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package envvar

import (
	"os"
	"testing"
)

func TestGetWithDefault(t *testing.T) {
	envVar := "TESTENVVAR_GETWITHDEFAULT"

	t.Run("missing", func(t *testing.T) { //nolint:paralleltest
		want := "default"

		os.Unsetenv(envVar)

		got := GetWithDefault(envVar, want)

		if got != want {
			t.Fatalf("expected %s, got %s", want, got)
		}
	})

	t.Run("empty", func(t *testing.T) {
		want := "default"

		t.Setenv(envVar, "")

		got := GetWithDefault(envVar, want)

		if got != want {
			t.Fatalf("expected %s, got %s", want, got)
		}
	})

	t.Run("not empty", func(t *testing.T) {
		want := "notempty"

		t.Setenv(envVar, want)

		got := GetWithDefault(envVar, "default")

		if got != want {
			t.Fatalf("expected %s, got %s", want, got)
		}
	})
}
