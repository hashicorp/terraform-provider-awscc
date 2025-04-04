# Create source location first
resource "awscc_mediatailor_source_location" "example" {
  source_location_name = "example-source-location"
  http_configuration = {
    base_url = "https://example.com/video"
  }

  tags = [{
    key   = "Modified By"
    value = "AWSCC"
  }]
}

# Example of awscc_mediatailor_live_source
resource "awscc_mediatailor_live_source" "example" {
  live_source_name     = "example-live-source"
  source_location_name = awscc_mediatailor_source_location.example.source_location_name

  http_package_configurations = [
    {
      path         = "/example/path"
      source_group = "example-source-group"
      type         = "HLS"
    }
  ]

  tags = [{
    key   = "Modified By"
    value = "AWSCC"
  }]
}