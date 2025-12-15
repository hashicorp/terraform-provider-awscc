terraform {
  provider_meta "awscc" {
    user_agent = [
      "example-demo/0.0.1 (a demo module)",
    ]
  }
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

