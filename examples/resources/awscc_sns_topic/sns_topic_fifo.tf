# The following resource creates an SNS First-In-First-Out (FIFO) Topic:
# Note: FIFO topic names must end with .fifo
resource "awscc_sns_topic" "sns_fifo_example" {
  topic_name                  = "sns-example.fifo"
  fifo_topic                  = true
  content_based_deduplication = true

  tags = [{
    key   = "Modified By"
    value = "AWSCC"
  }]
}
