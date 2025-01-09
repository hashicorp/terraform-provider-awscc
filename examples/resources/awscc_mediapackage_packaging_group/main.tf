# Get the current AWS region
data "aws_region" "current" {}

# Get the current AWS account ID
data "aws_caller_identity" "current" {}

# Create the MediaPackage Packaging Group
resource "awscc_mediapackage_packaging_group" "example" {
  packaging_group_id = "example-packaging-group"

  tags = [
    {
      key   = "Environment"
      value = "Production"
    },
    {
      key   = "Modified By"
      value = "AWSCC"
    }
  ]
}