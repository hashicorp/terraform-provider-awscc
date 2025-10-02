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

