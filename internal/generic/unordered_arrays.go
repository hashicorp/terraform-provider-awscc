// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package generic

import (
	"github.com/hashicorp/terraform-plugin-framework/resource/schema"
)

// unorderedArrayPaths walks a resource's Terraform schema and returns the set of
// CloudFormation JSON Pointer paths (array indices stripped, e.g. "/Config" or
// "/Foo/Bar") that back a Set-typed attribute (SetAttribute or SetNestedAttribute).
//
// The code generator already maps any CloudFormation array property declared
// "insertionOrder": false to a Set attribute, so an attribute's Set-ness is itself
// the authoritative signal that CloudFormation may return that array in a different
// order than it was submitted. Reusing that existing signal here -- instead of
// sniffing element shape for a "Key"/"key" field -- lets patchDocument's
// full-array-replace escape hatch fire for every order-unstable array, not just
// Tag-shaped ones.
func unorderedArrayPaths(attrs map[string]schema.Attribute, tfToCfNameMap map[string]string) map[string]bool {
	paths := make(map[string]bool)
	collectUnorderedArrayPaths(attrs, tfToCfNameMap, "", paths)
	return paths
}

func collectUnorderedArrayPaths(attrs map[string]schema.Attribute, tfToCfNameMap map[string]string, prefix string, paths map[string]bool) {
	for name, attr := range attrs {
		cfName, ok := tfToCfNameMap[name]
		if !ok {
			cfName = name
		}
		path := prefix + "/" + cfName

		switch a := attr.(type) {
		case schema.SetAttribute:
			paths[path] = true
		case schema.SetNestedAttribute:
			paths[path] = true
			collectUnorderedArrayPaths(a.NestedObject.Attributes, tfToCfNameMap, path, paths)
		case schema.ListNestedAttribute:
			collectUnorderedArrayPaths(a.NestedObject.Attributes, tfToCfNameMap, path, paths)
		case schema.SingleNestedAttribute:
			collectUnorderedArrayPaths(a.Attributes, tfToCfNameMap, path, paths)
		}
	}
}
