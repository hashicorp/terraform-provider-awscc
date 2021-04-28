package generator

import (
	"testing"

	cfschema "github.com/hashicorp/aws-cloudformation-resource-schema-sdk-go"
)

func RootPropertySchema_Test(t *testing.T) {
	resourceSchema, err := cfschema.NewResourceJsonSchemaPath("aws-logs-loggroup.json")
	if err != nil {
		t.Errorf("error loading test schema: %s", err)
	}

	resource, err := resourceSchema.Resource()
	if err != nil {
		t.Errorf("error parsing test schema: %s", err)
	}

	if err := resource.Expand(); err != nil {
		t.Errorf("error expanding JSON Pointer references: %s", err)
	}

	for propertyName, property := range resource.Properties {
		t.Logf("[DEBUG] Found schema property (%s):\n%s", propertyName, property)
		t.Logf("[DEBUG] Generated schema for property (%s): %s", propertyName, RootPropertySchema(resource, propertyName))
	}
}
