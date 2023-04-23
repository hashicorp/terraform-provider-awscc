# Copyright (c) HashiCorp, Inc.
# SPDX-License-Identifier: MPL-2.0

resource "awscc_accessanalyzer_analyzer" "this" {
  analyzer_name = "example"
  type          = ACCOUNT
}