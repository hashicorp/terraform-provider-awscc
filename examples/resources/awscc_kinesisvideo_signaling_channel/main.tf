# Kinesis Video Signaling Channel
resource "awscc_kinesisvideo_signaling_channel" "example" {
  name                = "example-signaling-channel"
  type                = "SINGLE_MASTER"
  message_ttl_seconds = 60

  tags = [{
    key   = "Modified By"
    value = "AWSCC"
  }]
}

# Output the ARN of the signaling channel
output "signaling_channel_arn" {
  value = awscc_kinesisvideo_signaling_channel.example.arn
}