resource "awscc_deadline_farm" "example" {
  display_name = "Example Farm"
  description  = "Example Deadline Farm"

  tags = [{
    key   = "ModifiedBy"
    value = "AWSCC"
  }]
}

resource "awscc_deadline_queue" "example" {
  display_name = "Example Queue"
  farm_id      = awscc_deadline_farm.example.farm_id
}

resource "awscc_deadline_queue_environment" "example" {
  farm_id       = awscc_deadline_farm.example.farm_id
  queue_id      = awscc_deadline_queue.example.queue_id
  priority      = 50
  template_type = "JSON"
  template = jsonencode({
    specificationVersion = "environment-2023-09"
    environment = {
      name = "ExampleEnvironment"
      variables = {
        EXAMPLE_VAR = "example_value"
      }
    }
  })
}
