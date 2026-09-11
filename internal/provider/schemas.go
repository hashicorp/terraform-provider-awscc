// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

// Legacy code: drives the legacy schema generator (generators/schema/main.go),
// superseded by bigdiffer's discovery + `-update`/`-generate`/`-docs`
// pipeline. Kept as the deprecated make/go:generate fallback; see
// contributing/docs/generating-the-provider-with-bigdiffer.md#fallback-the-legacy-process
// and contributing/docs/removing-the-legacy-generation-process.md.

//go:generate go run generators/schema/main.go -config all_schemas.hcl -generated-code-root .. -import-path-root github.com/hashicorp/terraform-provider-awscc/internal -- resources.go singular_data_sources.go plural_data_sources.go import_examples_gen.json

package provider
