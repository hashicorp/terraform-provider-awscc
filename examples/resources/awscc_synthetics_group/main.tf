# Example of creating a Synthetics Group using AWSCC provider
resource "awscc_synthetics_group" "example" {
  name = "example-synthetics-group"

  # Optional tags
  tags = [{
    key   = "Modified By"
    value = "AWSCC"
  }]
}