# Get current AWS account ID
data "aws_caller_identity" "current" {}

# Example of HTTP URL destination
resource "awscc_iot_topic_rule_destination" "http_destination" {
  http_url_properties = {
    confirmation_url = "https://example.com/confirm-${data.aws_caller_identity.current.account_id}"
  }
  status = "ENABLED"
}