# Create the SNS topic 
resource "aws_sns_topic" "example" {
  name = "example-topic"
  tags = {
    Environment = "Production"
    Name        = "example-rule"
  }
}

# EventBridge Rule resource
resource "awscc_events_rule" "example" {
  name        = "example-rule"
  description = "Example EventBridge Rule managed by Terraform"

  # Schedule-based rule that runs every 10 minutes
  schedule_expression = "rate(10 minutes)"

  # Enable the rule
  state = "ENABLED"

  # Target configuration - sending to the SNS topic
  targets = [
    {
      id  = "example-target"
      arn = aws_sns_topic.example.arn
    }
  ]
}
