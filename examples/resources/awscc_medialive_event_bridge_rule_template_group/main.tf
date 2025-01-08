# Data sources for AWS account ID and region
data "aws_caller_identity" "current" {}
data "aws_region" "current" {}

# Create MediaLive EventBridge Rule Template Group
resource "awscc_medialive_event_bridge_rule_template_group" "example" {
  name        = "example-template-group"
  description = "Example EventBridge Rule Template Group for MediaLive"

  tags = [{
    key   = "Environment"
    value = "Development"
  }, {
    key   = "Modified_By"
    value = "AWSCC"
  }]
}