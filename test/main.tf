terraform {
  required_providers {
    cloudapi = {
      source = "hashicorp/aws-cloudapi"
    }
  }
}

provider "cloudapi" {
  region = "us-west-2"
}

resource "aws_logs_log_group" "test" {
  provider = cloudapi

  log_group_name    = "CloudAPI_testing"
  retention_in_days = 120
}