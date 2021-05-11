package provider

import (
	"testing"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

var (
	supportedSchema = map[string]*schema.Schema{
		"required_string": {
			Type:     schema.TypeString,
			Required: true,
		},

		"optional_int": {
			Type:     schema.TypeInt,
			Optional: true,
		},

		"optional_float": {
			Type:     schema.TypeFloat,
			Optional: true,
		},
	}

	unsupportedSchema = map[string]*schema.Schema{
		"required_bool": {
			Type:     schema.TypeBool,
			Required: true,
		},

		"optional_list": {
			Type:     schema.TypeList,
			Optional: true,
			Elem:     &schema.Schema{Type: schema.TypeString},
		},
	}
)

func Test_StateConverter_UnsupportedSchema(t *testing.T) {
	converter := &StateConverter{
		TfSchema: unsupportedSchema,
	}
	r := &schema.Resource{
		Schema: converter.TfSchema,
	}

	d := r.Data(nil)
	d.Set("required_bool", true)

	if _, err := converter.ToCloudFormation(d); err == nil {
		t.Error("unexpected success ToCloudFormation")
	}

	d = r.Data(nil)

	if err := converter.ToTerraform(`{"RequiredBool": false}`, d); err != nil {
		t.Errorf("unexpected failure ToTerraform: %s", err)
	}
}

func Test_StateConverter_SupportedSchema(t *testing.T) {
	converter := &StateConverter{
		TfSchema: supportedSchema,
	}
	r := &schema.Resource{
		Schema: converter.TfSchema,
	}

	d := r.Data(nil)
	d.Set("required_string", "This is set")
	d.Set("optional_int", 42)

	s, err := converter.ToCloudFormation(d)

	if err != nil {
		t.Errorf("unexpected failure ToCloudFormation: %s", err)
	}
	if s == "" {
		t.Error("empty CloudFormation state")
	}

	//t.Log(s)

	d = r.Data(nil)

	if err := converter.ToTerraform(`{"RequiredString": "New value", "OptionalFloat": 4.2}`, d); err != nil {
		t.Errorf("unexpected failure ToTerraform: %s", err)
	}

	v := d.Get("required_string")
	s, ok := v.(string)
	if !ok {
		t.Errorf("unexpected type for required_string: %t", v)
	}
	if s != "New value" {
		t.Errorf("unexpected value for required_string: %s", s)
	}

	v = d.Get("optional_float")
	f, ok := v.(float64)
	if !ok {
		t.Errorf("unexpected type for optional_float: %t", v)
	}
	if f != 4.2 {
		t.Errorf("unexpected value for optional_float: %f", f)
	}
}
