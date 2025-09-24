# Create a Kinesis stream first
resource "aws_kinesis_stream" "example_stream" {
  name             = "example-stream"
  shard_count      = 1
  retention_period = 24

  tags = {
    Name        = "example-stream"
    Environment = "example"
  }
}

# Create the Kinesis stream consumer using AWS CloudControl
resource "awscc_kinesis_stream_consumer" "example_consumer" {
  consumer_name = "example-consumer"
  stream_arn    = aws_kinesis_stream.example_stream.arn

  tags = [
    {
      key   = "Name"
      value = "example-consumer"
    },
    {
      key   = "Environment"
      value = "example"
    }
  ]
}
