# Get current account ID
data "aws_caller_identity" "current" {}

# Get current region
data "aws_region" "current" {}

resource "awscc_connect_predefined_attribute" "example" {
  name         = "ExamplePredefinedAttribute"
  instance_arn = "arn:aws:connect:${data.aws_region.current.name}:${data.aws_caller_identity.current.account_id}:instance/example-instance-id"
  values = {
    string_list = ["Value1", "Value2"]
  }
}