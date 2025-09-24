resource "awscc_gamelift_location" "example" {
  location_name = "custom-example"

  tags = [{
    key   = "Modified By"
    value = "AWSCC"
  }]
}