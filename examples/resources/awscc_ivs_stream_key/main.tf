# Get current AWS region and account details
data "aws_region" "current" {}
data "aws_caller_identity" "current" {}

# First create an IVS Channel
resource "awscc_ivs_channel" "example" {
  name = "example-channel"
  tags = [{
    key   = "Modified By"
    value = "AWSCC"
  }]
}

# Create the IVS Stream Key
resource "awscc_ivs_stream_key" "example" {
  channel_arn = awscc_ivs_channel.example.arn
  tags = [{
    key   = "Modified By"
    value = "AWSCC"
  }]
}