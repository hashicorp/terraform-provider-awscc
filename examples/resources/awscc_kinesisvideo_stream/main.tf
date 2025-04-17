# Example configuration of an AWS Kinesis Video Stream
resource "awscc_kinesisvideo_stream" "example" {
  name                    = "example-video-stream"
  data_retention_in_hours = 24
  device_name             = "example-device"
  media_type              = "video/h264"
  tags = [{
    key   = "Modified By"
    value = "AWSCC"
  }]
}