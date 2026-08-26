// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package generic

import (
	"github.com/hashicorp/terraform-plugin-framework/resource/schema"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/planmodifier"
)

// unorderedArrayPaths walks a resource's Terraform schema and returns the set of
// CloudFormation JSON Pointer paths (array indices stripped, e.g. "/Config" or
// "/Foo/Bar") that back an order-unstable array attribute.
//
// A CloudFormation array property declared "insertionOrder": false may be returned
// by CloudFormation in a different order than it was submitted, so positional JSON
// Patch operations against it can corrupt data. The code generator surfaces such
// properties in one of two shapes, both of which are collected here:
//
//   - "insertionOrder": false + "uniqueItems": true  -> a Set attribute
//     (SetAttribute / SetNestedAttribute).
//   - "insertionOrder": false + non-unique items     -> a List attribute
//     (ListAttribute / ListNestedAttribute) carrying the generic.Multiset() plan
//     modifier, which is itself the authoritative marker of an unordered, non-unique
//     array.
//
// Reusing these existing schema signals -- instead of sniffing element shape for a
// "Key"/"key" field or re-deriving insertionOrder from the CFN schema at runtime --
// lets patchDocument's full-array-replace escape hatch fire for every order-unstable
// array, not just Tag-shaped ones.
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
		case schema.ListAttribute:
			// A List carrying the Multiset plan modifier is an "insertionOrder": false
			// array of non-unique items -- order-unstable, just like a Set.
			if hasMultisetPlanModifier(a.PlanModifiers) {
				paths[path] = true
			}
		case schema.ListNestedAttribute:
			if hasMultisetPlanModifier(a.PlanModifiers) {
				paths[path] = true
			}
			collectUnorderedArrayPaths(a.NestedObject.Attributes, tfToCfNameMap, path, paths)
		case schema.SingleNestedAttribute:
			collectUnorderedArrayPaths(a.Attributes, tfToCfNameMap, path, paths)
		}
	}
}

// hasMultisetPlanModifier reports whether the given list plan modifiers include the
// generic.Multiset() modifier, which the generator attaches to every List attribute
// derived from an "insertionOrder": false, non-unique CloudFormation array.
func hasMultisetPlanModifier(planModifiers []planmodifier.List) bool {
	for _, planModifier := range planModifiers {
		if _, ok := planModifier.(multisetAttributePlanModifier); ok {
			return true
		}
	}
	return false
}
