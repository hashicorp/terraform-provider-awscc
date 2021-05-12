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
		v, ok := d.GetOk(key)

		if !ok {
			continue
		}

		v, err := jsonValue(v, s.Type)

		if err != nil {
			return "", fmt.Errorf("error getting JSON type for key %q: %w", key, err)
		}

		m[strcase.ToCamel(key)] = v
	}

	cfState, err := json.Marshal(m)

	if err != nil {
		return "", err
	}

	return string(cfState), nil
}

func (c *StateConverter) ToCloudFormationChanges(d *schema.ResourceData) (string, string, error) {
	mOld := map[string]interface{}{}
	mNew := map[string]interface{}{}

	for key, s := range c.TfSchema {
		if d.HasChange(key) {
			o, n := d.GetChange(key)

			o, err := jsonValue(o, s.Type)

			if err != nil {
				return "", "", fmt.Errorf("error getting JSON type for key %q: %w", key, err)
			}

			n, err = jsonValue(n, s.Type)

			if err != nil {
				return "", "", fmt.Errorf("error getting JSON type for key %q: %w", key, err)
			}

			mOld[strcase.ToCamel(key)] = o
			mNew[strcase.ToCamel(key)] = n
		} else {
			v, ok := d.GetOk(key)

			if !ok {
				continue
			}

			v, err := jsonValue(v, s.Type)

			if err != nil {
				return "", "", fmt.Errorf("error getting JSON type for key %q: %w", key, err)
			}

			mOld[strcase.ToCamel(key)] = v
			mNew[strcase.ToCamel(key)] = v
		}
	}

	cfStateOld, err := json.Marshal(mOld)

	if err != nil {
		return "", "", err
	}

	cfStateNew, err := json.Marshal(mNew)

	if err != nil {
		return "", "", err
	}

	return string(cfStateOld), string(cfStateNew), nil
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

		v, err := tfValue(v, s.Type)

		if err != nil {
			return fmt.Errorf("error getting Terraform type for key %q: %w", key, err)
		}

		if err := d.Set(key, v); err != nil {
			return fmt.Errorf("error setting %s: %w", key, err)
		}
	}

	return nil
}

func jsonValue(v interface{}, t schema.ValueType) (interface{}, error) {
	switch t {
	case schema.TypeBool:
	case schema.TypeFloat:
	case schema.TypeInt:
		v = float64(v.(int))
	case schema.TypeString:
	default:
		return nil, fmt.Errorf("unsupported type (%T)", v)
	}

	return v, nil
}

func tfValue(v interface{}, t schema.ValueType) (interface{}, error) {
	switch t {
	case schema.TypeBool:
	case schema.TypeFloat:
	case schema.TypeInt:
		v = int(v.(float64))
	case schema.TypeString:
	default:
		return nil, fmt.Errorf("unsupported type (%T)", v)
	}

	return v, nil
}
