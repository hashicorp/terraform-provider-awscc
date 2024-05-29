resource "awscc_location_api_key" "example" {
  key_name    = "example_key"
  description = "Example Location API key"
  no_expiry   = true
  restrictions = {
    allow_actions   = ["geo:GetMap*", "geo:GetPlace"]
    allow_resources = ["arn:aws:geo:${data.aws_region.current.name}:${data.aws_caller_identity.current.account_id}:map/ExampleMap*"]
  }
  tags = [{
    key   = "Modified_By"
    value = "AWSCC"
  }]
}

data "aws_caller_identity" "current" {}
data "aws_region" "current" {}
