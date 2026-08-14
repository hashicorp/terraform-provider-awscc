resource "awscc_bedrockmantle_project" "example" {
  name = "example-project"
  tags = [
    {
      key   = "Environment"
      value = "example"
    },
    {
      key   = "Name"
      value = "example-project"
    }
  ]
}