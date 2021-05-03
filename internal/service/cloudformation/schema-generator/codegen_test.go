package generator

import (
	"testing"

	cfschema "github.com/hashicorp/aws-cloudformation-resource-schema-sdk-go"
)

func TestRootPropertySchema(t *testing.T) {
	resourceSchema, err := cfschema.NewResourceJsonSchemaPath("testdata/aws_logs_log_group.cf-resource-schema.json")
	if err != nil {
		t.Fatalf("error loading test schema: %s", err)
	}

	resource, err := resourceSchema.Resource()
	if err != nil {
		t.Fatalf("error parsing test schema: %s", err)
	}

	if err := resource.Expand(); err != nil {
		t.Fatalf("error expanding JSON Pointer references: %s", err)
	}

	for propertyName, property := range resource.Properties {
		t.Logf("[DEBUG] Found schema property (%s):\n%s", propertyName, property)
		t.Logf("[DEBUG] Generated schema for property (%s): %s", propertyName, RootPropertySchema(resource, propertyName))
	}
}
