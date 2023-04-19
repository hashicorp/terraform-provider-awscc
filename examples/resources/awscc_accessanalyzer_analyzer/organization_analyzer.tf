# Copyright (c) HashiCorp, Inc.
# SPDX-License-Identifier: MPL-2.0

resource "aws_organizations_organization" "this" {
  aws_service_access_principals = ["access-analyzer.amazonaws.com"]
}

resource "awscc_accessanalyzer_analyzer" "this" {
  analyzer_name = "example"
  type          = "ORGANIZATION"
}