# Get current AWS account ID
data "aws_caller_identity" "current" {}

# Get current region
data "aws_region" "current" {}

# Create a Panorama package
resource "awscc_panorama_package" "example" {
  package_name = "example-package"

  tags = [{
    key   = "Modified By"
    value = "AWSCC"
  }]
}