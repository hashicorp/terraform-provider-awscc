resource "awscc_deadline_farm" "example" {
  display_name = "ExampleRenderFarm"
  description  = "Example Deadline Farm for demonstrating limit configuration"

  tags = [
    {
      key   = "ModifiedBy"
      value = "AWSCC"
    }
  ]
}

# Create a Deadline Limit for CPU usage
resource "awscc_deadline_limit" "example" {
  farm_id                 = awscc_deadline_farm.example.farm_id
  display_name            = "CPU Limit"
  description             = "CPU core usage limit for the render farm"
  amount_requirement_name = "amount.cpu"
  max_count               = 100
}
