// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package main

import (
	"fmt"
	"sort"

	"github.com/hashicorp/hcl/v2"
	"github.com/hashicorp/hcl/v2/hclwrite"
	"github.com/zclconf/go-cty/cty"
)

// boolAttrs are the attributes written as bare booleans; everything else is a
// quoted string.
var boolAttrs = map[string]bool{
	attrNonProvisionable:                     true,
	attrSuppressResource:                     true,
	attrSuppressSingular:                     true,
	"suppress_plural_data_source_generation": true,
}

// setBlockAttributes sets (adds or updates) attributes on the single
// resource_schema block in blockText, preserving comments and existing
// attributes. It uses hclwrite for a token-preserving edit rather than string
// surgery, so free-form "# Suppression Reason" comments and commented-out lines
// survive. Attributes are applied in sorted order for deterministic output.
func setBlockAttributes(blockText string, attrs map[string]string) (string, error) {
	if len(attrs) == 0 {
		return blockText, nil
	}

	f, diags := hclwrite.ParseConfig([]byte(blockText), "block.hcl", hcl.InitialPos)
	if diags.HasErrors() {
		return "", fmt.Errorf("parsing block: %s", diags.Error())
	}

	blocks := f.Body().Blocks()
	if len(blocks) != 1 {
		return "", fmt.Errorf("expected exactly one resource_schema block, got %d", len(blocks))
	}
	body := blocks[0].Body()

	for _, name := range sortedKeys(attrs) {
		body.SetAttributeValue(name, attrValue(name, attrs[name]))
	}

	return string(f.Bytes()), nil
}

// attrValue renders a policy attribute value as the correct cty type: a bare
// boolean for the suppress_*/non_provisionable flags, a quoted string otherwise
// (e.g. frozen_since dates).
func attrValue(name, val string) cty.Value {
	if boolAttrs[name] {
		return cty.BoolVal(val == "true")
	}
	return cty.StringVal(val)
}

func sortedKeys(m map[string]string) []string {
	keys := make([]string, 0, len(m))
	for k := range m {
		keys = append(keys, k)
	}
	sort.Strings(keys)
	return keys
}
