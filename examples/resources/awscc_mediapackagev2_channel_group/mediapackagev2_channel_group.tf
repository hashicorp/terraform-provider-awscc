resource "awscc_mediapackagev2_channel_group" "example" {
  channel_group_name = "example"
  description        = "example"

  tags = [{
    key   = "Modified By"
    value = "AWSCC"
  }]
}
