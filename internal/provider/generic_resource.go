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

// Create is the generic Create handler for a generated resource.
func (g *GenericResource) Create(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	instance := GenericResourceInstance{
		GenericResource: g,
		ResourceData:    d,
	}

	return instance.Create(ctx)
}

// Read is the generic Read handler for a generated resource.
func (g *GenericResource) Read(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	instance := GenericResourceInstance{
		GenericResource: g,
		ResourceData:    d,
	}

	return instance.Read(ctx)
}

// Update is the generic Update handler for a generated resource.
func (g *GenericResource) Update(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	instance := GenericResourceInstance{
		GenericResource: g,
		ResourceData:    d,
	}

	return instance.Update(ctx)
}

// Delete is the generic Delete handler for a generated resource.
func (g *GenericResource) Delete(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	instance := GenericResourceInstance{
		GenericResource: g,
		ResourceData:    d,
	}

	return instance.Delete(ctx)
}

type GenericResourceInstance struct {
	GenericResource *GenericResource
	ResourceData    *schema.ResourceData
}

// Create is the generic Create handler for a generated resource.
func (g *GenericResourceInstance) Create(ctx context.Context) diag.Diagnostics {
	return nil
}

// Read is the generic Read handler for a generated resource.
func (g *GenericResourceInstance) Read(ctx context.Context) diag.Diagnostics {
	return nil
}

// Update is the generic Update handler for a generated resource.
func (g *GenericResourceInstance) Update(ctx context.Context) diag.Diagnostics {
	return nil
}

// Delete is the generic Delete handler for a generated resource.
func (g *GenericResourceInstance) Delete(ctx context.Context) diag.Diagnostics {
	return nil
}
