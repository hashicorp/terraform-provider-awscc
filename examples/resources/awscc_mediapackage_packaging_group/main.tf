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