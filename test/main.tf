provider "aws-cloudapi" {
  region = "us-west-2"
}

resource "aws_logs_log_group" "test" {
    provider = "aws-cloudapi"

    log_group_name    = "CloudAPI_testing"
    retention_in_days = 120
}