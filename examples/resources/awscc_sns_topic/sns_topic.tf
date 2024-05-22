# The following resource creates an SNS Topic:
resource "awscc_sns_topic" "sns_example" {
  topic_name = "sns-example-topic"

  tags = [{
    key   = "Modified By"
    value = "AWSCC"
  }]
}