# Create a Panorama package
resource "awscc_panorama_package" "example" {
  package_name = "example-package"

  tags = [{
    key   = "Modified By"
    value = "AWSCC"
  }]
}