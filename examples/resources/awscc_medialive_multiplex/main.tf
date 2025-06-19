data "aws_caller_identity" "current" {}
# Note: Using data.aws_region.current.region (AWS provider v6.0+)
# For AWS provider < v6.0, use data.aws_region.current.name instead
data "aws_region" "current" {}

# Create the MediaLive Multiplex
resource "awscc_medialive_multiplex" "example" {
  name               = "example-multiplex"
  availability_zones = ["${data.aws_region.current.region}a", "${data.aws_region.current.region}b"]

  multiplex_settings = {
    transport_stream_bitrate                = 1000000
    transport_stream_id                     = 1
    maximum_video_buffer_delay_milliseconds = 1000
  }

  destinations = [{
    # Using media connect output destination
    multiplex_media_connect_output_destination_settings = {
      entitlement_arn = "arn:aws:mediaconnect:${data.aws_region.current.region}:${data.aws_caller_identity.current.account_id}:entitlement:1234-5678-90ab-cdef"
    }
  }]

  tags = [{
    key   = "Modified By"
    value = "AWSCC"
  }]
}