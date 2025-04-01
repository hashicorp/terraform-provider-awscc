# First create a channel group
resource "awscc_mediapackagev2_channel_group" "example" {
  channel_group_name = "example-channel-group"

  tags = [{
    key   = "Modified By"
    value = "AWSCC"
  }]
}

# Then create a channel
resource "awscc_mediapackagev2_channel" "example" {
  channel_group_name = awscc_mediapackagev2_channel_group.example.channel_group_name
  channel_name       = "example-channel"

  tags = [{
    key   = "Modified By"
    value = "AWSCC"
  }]
}

# Finally create the origin endpoint
resource "awscc_mediapackagev2_origin_endpoint" "example" {
  channel_group_name   = awscc_mediapackagev2_channel_group.example.channel_group_name
  channel_name         = awscc_mediapackagev2_channel.example.channel_name
  origin_endpoint_name = "example-endpoint"
  container_type       = "TS"

  # Optional HLS manifest configuration
  hls_manifests = [{
    manifest_name                      = "index"
    manifest_window_seconds            = 60
    program_date_time_interval_seconds = 10
  }]

  # Optional segment configuration
  segment = {
    segment_duration_seconds    = 6
    segment_name                = "seg"
    include_iframe_only_streams = true
  }

  tags = [{
    key   = "Modified By"
    value = "AWSCC"
  }]
}