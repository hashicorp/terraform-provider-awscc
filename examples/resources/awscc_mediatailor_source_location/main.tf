# Get current AWS region
data "aws_region" "current" {}

# AWS MediaTailor Source Location example
resource "awscc_mediatailor_source_location" "example" {
  source_location_name = "example-source-location"
  http_configuration = {
    base_url = "https://s3.${data.aws_region.current.name}.amazonaws.com/example-bucket/content"
  }

  tags = [{
    key   = "Modified By"
    value = "AWSCC"
  }]
}