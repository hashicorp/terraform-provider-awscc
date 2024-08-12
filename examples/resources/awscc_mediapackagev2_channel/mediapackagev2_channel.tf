resource "awscc_mediapackagev2_channel" "example" {
  channel_group_name = awscc_mediapackagev2_channel_group.example.channel_group_name
  channel_name       = "example"
  description        = "example"
  input_type         = "HLS"

  tags = [{
    key   = "Modified By"
    value = "AWSCC"
  }]
}

resource "awscc_mediapackagev2_channel_group" "example" {
  channel_group_name = "example"
  description        = "example"

  tags = [{
    key   = "Modified By"
    value = "AWSCC"
  }]
}
