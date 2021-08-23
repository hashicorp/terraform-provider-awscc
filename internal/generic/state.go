package generic

import (
	"context"

	"github.com/hashicorp/terraform-plugin-framework/tfsdk"
	"github.com/hashicorp/terraform-plugin-go/tftypes"
	"github.com/hashicorp/terraform-provider-awscc/internal/tfresource"
)

// CopyValueAtPath copies the value at a specified path from source State to destination State.
func CopyValueAtPath(ctx context.Context, dst, src *tfsdk.State, path *tftypes.AttributePath) error {
	val, diags := src.GetAttribute(ctx, path)

	if tfresource.DiagsHasError(diags) {
		return tfresource.DiagsError(diags)
	}

	diags = dst.SetAttribute(ctx, path, val)

	if tfresource.DiagsHasError(diags) {
		return tfresource.DiagsError(diags)
	}

	return nil
}
