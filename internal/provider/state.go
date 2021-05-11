package provider

import (
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

// Completely naive Terraform "state" to CloudFormation state conversion.

type StateConverter struct {
	TfSchema map[string]*schema.Schema
}

func NewStateConverter(schema map[string]*schema.Schema) *StateConverter {
	return &StateConverter{
		TfSchema: schema,
	}
}

func (c *StateConverter) ToCloudFormation(d *schema.ResourceData) (string, error) {
	return "", nil
}

func (c *StateConverter) ToTerraform(cfState string, d *schema.ResourceData) error {
	return nil
}
