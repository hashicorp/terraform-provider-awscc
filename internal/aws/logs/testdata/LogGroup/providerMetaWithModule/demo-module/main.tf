terraform {
  # custom provider_meta which should appear appended to user agent
  provider_meta "awscc" {
    user_agent = [
      {
        product_name    = "example-demo"
        product_version = "0.0.1"
        comment         = "a demo module"
      },
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
