// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package codegen

import (
	"go/format"
	"regexp"
	"strings"
	"testing"

	cfschema "github.com/hashicorp/aws-cloudformation-resource-schema-sdk-go"
	"github.com/hashicorp/cli"
)

func TestEmitterAttributeFunctionDeduplication(t *testing.T) {
	t.Parallel()

	testCases := map[string]struct {
		createOnlyProperties cfschema.PropertyJsonPointers
		expectedFunctions    int
		expectedSameRootCall bool
	}{
		"identical rendered attributes": {
			expectedFunctions:    2,
			expectedSameRootCall: true,
		},
		"path-specific attributes": {
			createOnlyProperties: cfschema.PropertyJsonPointers{
				cfschema.PropertyJsonPointer("/properties/First"),
			},
			expectedFunctions:    3,
			expectedSameRootCall: false,
		},
	}

	for name, testCase := range testCases {
		t.Run(name, func(t *testing.T) {
			t.Parallel()

			resource := testNestedResource()
			resource.CreateOnlyProperties = testCase.createOnlyProperties

			rootSchema, attributeFunctions := emitTestResource(t, resource, false)
			rootCalls := regexp.MustCompile(`"(?:first|second)":(schemaAttribute[0-9a-f]+)\(\)`).FindAllStringSubmatch(rootSchema, -1)
			if got, want := len(rootCalls), 2; got != want {
				t.Fatalf("root attribute call count = %d, want %d\n%s", got, want, rootSchema)
			}

			callsSameFunction := rootCalls[0][1] == rootCalls[1][1]
			if got, want := callsSameFunction, testCase.expectedSameRootCall; got != want {
				t.Fatalf("root attributes call same function = %t, want %t\n%s", got, want, rootSchema)
			}

			if got, want := strings.Count(attributeFunctions, "func schemaAttribute"), testCase.expectedFunctions; got != want {
				t.Fatalf("attribute function count = %d, want %d\n%s", got, want, attributeFunctions)
			}

			source := "package generated\n\n" + attributeFunctions + "\nvar _ = " + rootSchema
			if _, err := format.Source([]byte(source)); err != nil {
				t.Fatalf("generated code is not valid Go: %s\n%s", err, source)
			}
		})
	}
}

func TestEmitterAttributeFunctionsAreDeterministic(t *testing.T) {
	t.Parallel()

	firstRootSchema, firstAttributeFunctions := emitTestResource(t, testNestedResource(), false)
	secondRootSchema, secondAttributeFunctions := emitTestResource(t, testNestedResource(), false)

	if firstRootSchema != secondRootSchema {
		t.Fatalf("root schema changed between runs:\nfirst:\n%s\nsecond:\n%s", firstRootSchema, secondRootSchema)
	}
	if firstAttributeFunctions != secondAttributeFunctions {
		t.Fatalf("attribute functions changed between runs:\nfirst:\n%s\nsecond:\n%s", firstAttributeFunctions, secondAttributeFunctions)
	}
}

func TestAttributeFunctionNamesAreScoped(t *testing.T) {
	t.Parallel()

	body := "schema.StringAttribute{}"
	resourceRegistry := newAttributeFunctionRegistry("awscc_test_thing", false)
	dataSourceRegistry := newAttributeFunctionRegistry("awscc_test_thing", true)
	otherResourceRegistry := newAttributeFunctionRegistry("awscc_test_other", false)

	resourceName, err := resourceRegistry.register(body)
	if err != nil {
		t.Fatal(err)
	}
	dataSourceName, err := dataSourceRegistry.register(body)
	if err != nil {
		t.Fatal(err)
	}
	otherResourceName, err := otherResourceRegistry.register(body)
	if err != nil {
		t.Fatal(err)
	}

	if resourceName == dataSourceName || resourceName == otherResourceName || dataSourceName == otherResourceName {
		t.Fatalf("attribute function names are not scoped: resource=%q data_source=%q other_resource=%q", resourceName, dataSourceName, otherResourceName)
	}
}

func TestAttributeFunctionRenderHandlesLeadingComment(t *testing.T) {
	t.Parallel()

	registry := newAttributeFunctionRegistry("awscc_test_thing", false)
	body := "// Pattern: \"\"\nschema.MapAttribute{}"
	if _, err := registry.register(body); err != nil {
		t.Fatal(err)
	}

	rendered := registry.render()
	if want := "return (\n" + body + ")"; !strings.Contains(rendered, want) {
		t.Fatalf("attribute function does not parenthesize its body:\n%s", rendered)
	}

	source := "package generated\n\n" + rendered
	if _, err := format.Source([]byte(source)); err != nil {
		t.Fatalf("generated code is not valid Go: %s\n%s", err, source)
	}
}

func emitTestResource(t *testing.T, resource *cfschema.Resource, isDataSource bool) (string, string) {
	t.Helper()

	var rootSchema strings.Builder
	emitter := Emitter{
		CfResource:   resource,
		IsDataSource: isDataSource,
		Ui:           cli.NewMockUi(),
		Writer:       &rootSchema,
		Deduplicate:  true,
	}

	_, attributeFunctions, err := emitter.EmitRootPropertiesSchema("awscc_test_thing", make(map[string]string))
	if err != nil {
		t.Fatal(err)
	}

	return rootSchema.String(), attributeFunctions
}

func testNestedResource() *cfschema.Resource {
	stringType := cfschema.Type(cfschema.PropertyTypeString)
	objectType := cfschema.Type(cfschema.PropertyTypeObject)
	nestedProperty := func() *cfschema.Property {
		return &cfschema.Property{
			Type: &objectType,
			Properties: map[string]*cfschema.Property{
				"Value": {
					Type: &stringType,
				},
			},
		}
	}

	return &cfschema.Resource{
		Properties: map[string]*cfschema.Property{
			"First":  nestedProperty(),
			"Second": nestedProperty(),
		},
	}
}
