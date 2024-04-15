// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package defaults

import (
	"context"
	"testing"

	"github.com/hashicorp/terraform-plugin-framework/resource/schema/defaults"
)

func TestStaticSetOfString(t *testing.T) {
	t.Parallel()

	ctx := context.TODO()
	request := defaults.SetRequest{}
	response := defaults.SetResponse{}
	StaticSetOfString("One", "bee").DefaultSet(ctx, request, &response)

	if response.Diagnostics.HasError() {
		t.Errorf("unexpected error: %v", response.Diagnostics)
	}
	if want, got := 2, len(response.PlanValue.Elements()); got != want {
		t.Errorf("StaticListOfString.PlanValue length %d, want %d", got, want)
	}
}
