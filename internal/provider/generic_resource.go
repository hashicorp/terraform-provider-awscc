package provider

import (
	"context"

	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

type GenericResource struct {
	CloudFormationTypeName string
	TerraformSchema        map[string]*schema.Schema
	TerraformTypeName      string
}

// GetSchema returns the resource's schema (currently the schema.Resource)
func (g *GenericResource) GetSchema() *schema.Resource {
	return &schema.Resource{
		CreateContext: g.Create,
		ReadContext:   g.Read,
		UpdateContext: g.Update,
		DeleteContext: g.Delete,

		Schema: g.TerraformSchema,
	}
}

func (g *GenericResource) ResourceInstance(ctx context.Context, meta interface{}) *GenericResourceInstance {
	return &GenericResourceInstance{
		GenericResource: g,
		Meta:            meta,
	}
}

// Create is the generic Create handler for a generated resource.
func (g *GenericResource) Create(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	return g.ResourceInstance(ctx, meta).Create(ctx, d)
}

// Read is the generic Read handler for a generated resource.
func (g *GenericResource) Read(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	return g.ResourceInstance(ctx, meta).Read(ctx, d)
}

// Update is the generic Update handler for a generated resource.
func (g *GenericResource) Update(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	return g.ResourceInstance(ctx, meta).Update(ctx, d)
}

// Delete is the generic Delete handler for a generated resource.
func (g *GenericResource) Delete(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	return g.ResourceInstance(ctx, meta).Delete(ctx, d)
}

type GenericResourceInstance struct {
	GenericResource *GenericResource
	Meta            interface{}
}

// Create is the generic Create handler for a generated resource.
func (g *GenericResourceInstance) Create(ctx context.Context, d *schema.ResourceData) diag.Diagnostics {
	return nil
}

// Read is the generic Read handler for a generated resource.
func (g *GenericResourceInstance) Read(ctx context.Context, d *schema.ResourceData) diag.Diagnostics {
	return nil
}

// Update is the generic Update handler for a generated resource.
func (g *GenericResourceInstance) Update(ctx context.Context, d *schema.ResourceData) diag.Diagnostics {
	return nil
}

// Delete is the generic Delete handler for a generated resource.
func (g *GenericResourceInstance) Delete(ctx context.Context, d *schema.ResourceData) diag.Diagnostics {
	return nil
}
