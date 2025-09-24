// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

//go:generate go run generators/schema/main.go -config all_schemas.hcl -generated-code-root .. -import-path-root github.com/hashicorp/terraform-provider-awscc/internal -- resources.go singular_data_sources.go plural_data_sources.go import_examples_gen.json

package provider
