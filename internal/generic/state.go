package generic

import (
	"context"

	"github.com/hashicorp/terraform-plugin-framework/attr"
	"github.com/hashicorp/terraform-plugin-framework/path"
	"github.com/hashicorp/terraform-plugin-framework/tfsdk"
	"github.com/hashicorp/terraform-provider-awscc/internal/tfresource"
)

// CopyValueAtPath copies the value at a specified path from source State to destination State.
func CopyValueAtPath(ctx context.Context, dst, src *tfsdk.State, p path.Path) error {
	var val attr.Value
	diags := src.GetAttribute(ctx, p, &val)

	if diags.HasError() {
		return tfresource.DiagsError(diags)
	}

	diags = dst.SetAttribute(ctx, p, val)

	if diags.HasError() {
		return tfresource.DiagsError(diags)
	}

	return nil
}
