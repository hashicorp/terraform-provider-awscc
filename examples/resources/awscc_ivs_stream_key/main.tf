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