# Get current region
data "aws_region" "current" {}

# Create the MediaTailor source location first (required parent resource)
resource "awscc_mediatailor_source_location" "example" {
  source_location_name = "example-source-location"
  http_configuration = {
    base_url = "https://example-vod-source.s3.${data.aws_region.current.name}.amazonaws.com/content/"
  }

  tags = [{
    key   = "Modified By"
    value = "AWSCC"
  }]
}

# Create the VOD source
resource "awscc_mediatailor_vod_source" "example" {
  source_location_name = awscc_mediatailor_source_location.example.source_location_name
  vod_source_name      = "example-vod-source"

  http_package_configurations = [
    {
      path         = "path/to/content/"
      source_group = "example-source-group"
      type         = "HLS"
    }
  ]

  tags = [{
    key   = "Modified By"
    value = "AWSCC"
  }]
}