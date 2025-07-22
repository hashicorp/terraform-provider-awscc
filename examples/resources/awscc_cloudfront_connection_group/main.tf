data "aws_caller_identity" "current" {}
data "aws_region" "current" {}

resource "awscc_cloudfront_connection_group" "example" {
  name          = "example-connection-group"
  enabled       = true
  ipv_6_enabled = true

  tags = [{
    key   = "Modified By"
    value = "AWSCC"
  }]
}