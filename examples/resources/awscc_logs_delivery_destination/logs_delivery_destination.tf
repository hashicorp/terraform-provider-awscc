# CloudWatch Logs log group to use as a destination
resource "aws_cloudwatch_log_group" "example_destination" {
  name              = "example-destination-log-group"
  retention_in_days = 14

  tags = {
    Name        = "example-destination-log-group"
    Environment = "example"
  }
}

# Logs Delivery Destination
resource "awscc_logs_delivery_destination" "example" {
  name                     = "example-delivery-destination"
  destination_resource_arn = aws_cloudwatch_log_group.example_destination.arn
  output_format            = "json"

  tags = [
    {
      key   = "Name"
      value = "example-delivery-destination"
    },
    {
      key   = "Environment"
      value = "example"
    }
  ]
}
