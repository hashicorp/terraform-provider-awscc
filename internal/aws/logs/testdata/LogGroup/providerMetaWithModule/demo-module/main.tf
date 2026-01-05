# Copyright IBM Corp. 2021, 2026
# SPDX-License-Identifier: MPL-2.0

terraform {
  # custom provider_meta which should appear appended to user agent
  provider_meta "awscc" {
    user_agent = [
      "example-demo/0.0.1 (a demo module)",
    ]
  }
}

resource "awscc_logs_log_group" "test_module" {
  log_group_name    = var.rName
  retention_in_days = 7
}

variable "rName" {
  description = "Name for resource"
  type        = string
  nullable    = false
}
