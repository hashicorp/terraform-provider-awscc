// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package main

// bigdiffer owns the definition of all_schemas.hcl's shape outright, rather than
// importing it from the generator packages. This keeps the tool self-contained
// so the legacy generator/update code can be deleted once bigdiffer stabilizes.

// resourceRow is a single resource_schema block. It is a superset covering both
// the curated overlay (all_schemas.hcl) blocks and the leaner blocks produced by
// live discovery, so one type serves for parsing, discovery, and reconciliation.
//
// Suppression and freeze reasons are per-fact, not one shared string: a
// resource, its singular data source, and its plural data source can each be
// suppressed independently (item 4, "per-artifact independence"), and a freeze
// is a fourth, orthogonal fact (it pins the whole schema, not one artifact) —
// so each gets its own reason field, mirroring the existing one-flag-per-
// artifact shape of the Suppress*Generation booleans rather than one field
// trying to describe more than one fact at once
// (contributing/docs/suppressed-and-frozen.md).
type resourceRow struct {
	ResourceTypeName                     string `hcl:"resource_type_name,label"` // Terraform type name (block label)
	CloudFormationSchemaPath             string `hcl:"cloudformation_schema_path,optional"`
	CloudFormationTypeName               string `hcl:"cloudformation_type_name"`
	PathAwareAttributeNames              bool   `hcl:"path_aware_attribute_names,optional"`
	SuppressionReasonResource            string `hcl:"suppression_reason_resource,optional"`
	SuppressionReasonSingularDataSource  string `hcl:"suppression_reason_singular_data_source,optional"`
	SuppressionReasonPluralDataSource    string `hcl:"suppression_reason_plural_data_source,optional"`
	SuppressPluralDataSourceGeneration   bool   `hcl:"suppress_plural_data_source_generation,optional"`
	SuppressResourceGeneration           bool   `hcl:"suppress_resource_generation,optional"`
	SuppressSingularDataSourceGeneration bool   `hcl:"suppress_singular_data_source_generation,optional"`
	FrozenSince                          string `hcl:"frozen_since,optional"`
	FrozenReason                         string `hcl:"frozen_reason,optional"`
	NonProvisionable                     bool   `hcl:"non_provisionable,optional"`
}

// allSchemasFile mirrors the whole of all_schemas.hcl for a structural decode.
type allSchemasFile struct {
	Defaults  defaultsBlock   `hcl:"defaults,block"`
	Meta      metaSchemaBlock `hcl:"meta_schema,block"`
	Resources []resourceRow   `hcl:"resource_schema,block"`
}

// availableFile mirrors a bare list of resource_schema blocks (no defaults or
// meta_schema). It is a convenience decode target for tests that parse a single
// block or a fragment.
type availableFile struct {
	Resources []resourceRow `hcl:"resource_schema,block"`
}

type defaultsBlock struct {
	SchemaCacheDirectory    string `hcl:"schema_cache_directory"`
	TerraformTypeNamePrefix string `hcl:"terraform_type_name_prefix,optional"`
}

type metaSchemaBlock struct {
	Path string `hcl:"path"`
}
