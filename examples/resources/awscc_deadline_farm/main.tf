# Create the Deadline Farm
resource "awscc_deadline_farm" "example" {
  display_name = "ExampleFarm"
  description  = "Example Deadline Farm created with AWSCC provider"

  tags = [{
    key   = "Modified By"
    value = "AWSCC"
  }]
}