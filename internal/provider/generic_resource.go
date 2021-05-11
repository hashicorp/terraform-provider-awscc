package provider

import (
	"context"
	"fmt"
	"log"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/service/cloudformation"
	"github.com/hashicorp/aws-sdk-go-base/tfawserr"
	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/resource"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"github.com/hashicorp/terraform-provider-aws-cloudapi/internal/service/cloudformation/waiter"
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
		StateConverter:  NewStateConverter(g.TerraformSchema),
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
	StateConverter  *StateConverter
}

// Create is the generic Create handler for a generated resource.
func (g *GenericResourceInstance) Create(ctx context.Context, d *schema.ResourceData) diag.Diagnostics {
	conn := g.Meta.(*AWSClient).cfconn

	cfTypeName := g.GenericResource.CloudFormationTypeName
	tfTypeName := g.GenericResource.TerraformTypeName
	desiredState, err := g.StateConverter.ToCloudFormation(d)

	if err != nil {
		return diag.FromErr(fmt.Errorf("error converting Terraform state (%s/%s): %w", tfTypeName, cfTypeName, err))
	}

	input := &cloudformation.CreateResourceInput{
		ClientToken:  aws.String(resource.UniqueId()),
		DesiredState: aws.String(desiredState),
		TypeName:     aws.String(cfTypeName),
	}

	output, err := conn.CreateResourceWithContext(ctx, input)

	if err != nil {
		return diag.FromErr(fmt.Errorf("error creating CloudFormation Resource (%s/%s): %w", tfTypeName, cfTypeName, err))
	}

	if output == nil || output.ProgressEvent == nil {
		return diag.FromErr(fmt.Errorf("error creating CloudFormation Resource (%s/%s): empty_response", tfTypeName, cfTypeName))
	}

	// Always try to capture the identifier before returning errors
	d.SetId(aws.StringValue(output.ProgressEvent.Identifier))

	output.ProgressEvent, err = waiter.ResourceRequestStatusProgressEventOperationStatusSuccess(ctx, conn, aws.StringValue(output.ProgressEvent.RequestToken), d.Timeout(schema.TimeoutCreate))

	if err != nil {
		return diag.FromErr(fmt.Errorf("error waiting for CloudFormation Resource (%s) creation: %w", d.Id(), err))
	}

	// Some resources do not set the identifier until after creation
	if d.Id() == "" {
		d.SetId(aws.StringValue(output.ProgressEvent.Identifier))
	}

	return g.Read(ctx, d)
}

// Read is the generic Read handler for a generated resource.
func (g *GenericResourceInstance) Read(ctx context.Context, d *schema.ResourceData) diag.Diagnostics {
	conn := g.Meta.(*AWSClient).cfconn

	input := &cloudformation.GetResourceInput{
		Identifier: aws.String(d.Id()),
		TypeName:   aws.String(g.GenericResource.CloudFormationTypeName),
	}

	output, err := conn.GetResourceWithContext(ctx, input)

	if tfawserr.ErrCodeEquals(err, cloudformation.ErrCodeResourceNotFoundException) {
		if d.IsNewResource() {
			return diag.FromErr(fmt.Errorf("error reading CloudFormation Resource (%s): not found after creation", d.Id()))
		}

		log.Printf("[WARN] CloudFormation Resource (%s) not found, removing from state", d.Id())
		d.SetId("")
		return nil
	}

	if err != nil {
		return diag.FromErr(fmt.Errorf("error reading CloudFormation Resource (%s): %w", d.Id(), err))
	}

	if output == nil || output.ResourceDescription == nil {
		return diag.FromErr(fmt.Errorf("error reading CloudFormation Resource (%s): empty response", d.Id()))
	}

	if err := g.StateConverter.ToTerraform(aws.StringValue(output.ResourceDescription.ResourceModel), d); err != nil {
		return diag.FromErr(fmt.Errorf("error converting CloudFormation state (%s/%s): %w", g.GenericResource.TerraformTypeName, g.GenericResource.CloudFormationTypeName, err))
	}

	return nil
}

// Update is the generic Update handler for a generated resource.
func (g *GenericResourceInstance) Update(ctx context.Context, d *schema.ResourceData) diag.Diagnostics {
	return g.Read(ctx, d)
}

// Delete is the generic Delete handler for a generated resource.
func (g *GenericResourceInstance) Delete(ctx context.Context, d *schema.ResourceData) diag.Diagnostics {
	conn := g.Meta.(*AWSClient).cfconn

	input := &cloudformation.DeleteResourceInput{
		ClientToken: aws.String(resource.UniqueId()),
		Identifier:  aws.String(d.Id()),
	}

	output, err := conn.DeleteResourceWithContext(ctx, input)

	if err != nil {
		return diag.FromErr(fmt.Errorf("error deleting CloudFormation Resource (%s): %w", d.Id(), err))
	}

	if output == nil || output.ProgressEvent == nil {
		return diag.FromErr(fmt.Errorf("error deleting CloudFormation Resource (%s): empty response", d.Id()))
	}

	progressEvent, err := waiter.ResourceRequestStatusProgressEventOperationStatusSuccess(ctx, conn, aws.StringValue(output.ProgressEvent.RequestToken), d.Timeout(schema.TimeoutDelete))

	if progressEvent != nil && aws.StringValue(progressEvent.ErrorCode) == cloudformation.HandlerErrorCodeNotFound {
		return nil
	}

	if err != nil {
		return diag.FromErr(fmt.Errorf("error waiting for CloudFormation Resource (%s) deletion: %w", d.Id(), err))
	}

	return nil
}
