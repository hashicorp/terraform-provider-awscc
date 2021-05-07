package provider

import (
	"context"

	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

type GeneratedResource struct {
	CloudFormationTypeName string
	TerraformTypeName      string
}

// Create is the generic Create handler for a generated resource.
func (g *GeneratedResource) Create(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	return nil
}

// Read is the generic Read handler for a generated resource.
func (g *GeneratedResource) Read(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	return nil
}

// Update is the generic Update handler for a generated resource.
func (g *GeneratedResource) Update(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	return nil
}

// Delete is the generic Delete handler for a generated resource.
func (g *GeneratedResource) Delete(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	return nil
}
