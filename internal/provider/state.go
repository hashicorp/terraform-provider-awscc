package provider

import (
	"encoding/json"
	"fmt"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"github.com/iancoleman/strcase"
)

// Completely naive Terraform "state" to CloudFormation state conversion.

type StateConverter struct {
	TfSchema map[string]*schema.Schema
}

func (c *StateConverter) ToCloudFormation(d *schema.ResourceData) (string, error) {
	m := map[string]interface{}{}

	for key, s := range c.TfSchema {
		v := d.Get(key)

		switch s.Type {
		case schema.TypeBool:
		case schema.TypeFloat:
		case schema.TypeInt:
			v = float64(v.(int))
		case schema.TypeString:
		default:
			return "", fmt.Errorf("unsupported type (%t) for key %q", v, key)
		}

		if !s.Computed {
			m[strcase.ToCamel(key)] = v
		}
	}

	cfState, err := json.Marshal(m)

	if err != nil {
		return "", nil
	}

	return string(cfState), nil
}

func (c *StateConverter) ToTerraform(cfState string, d *schema.ResourceData) error {
	var v interface{}

	if err := json.Unmarshal([]byte(cfState), &v); err != nil {
		return err
	}

	m := v.(map[string]interface{})

	for key, s := range c.TfSchema {
		v, ok := m[strcase.ToCamel(key)]

		if !ok {
			continue
		}

		switch s.Type {
		case schema.TypeBool:
		case schema.TypeFloat:
		case schema.TypeInt:
			v = int(v.(float64))
		case schema.TypeString:
		default:
			return fmt.Errorf("unsupported type (%t) for key %q", v, key)
		}

		if err := d.Set(key, v); err != nil {
			return fmt.Errorf("error setting %s: %w", key, err)
		}
	}

	return nil
}
