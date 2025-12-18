# Copyright IBM Corp. 2021, 2025
# SPDX-License-Identifier: MPL-2.0

module "demo-module" {
  source = "./demo-module"
  rName  = var.rName
}

resource "awscc_logs_log_group" "test" {
  log_group_name    = var.rName
  retention_in_days = 7
}

variable "rName" {
  description = "Name for resource"
  type        = string
  nullable    = false
}

