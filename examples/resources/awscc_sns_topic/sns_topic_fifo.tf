# The following resource creates a SNS First-In-First-Out (FIFO) Topic:
# Note: FIFO topic names must end with .fifo
resource "aws_sns_topic" "sns_fifo_example" {
  name                        = "sns-example.fifo"
  fifo_topic                  = true
  content_based_deduplication = true
}