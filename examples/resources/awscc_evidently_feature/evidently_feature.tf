resource "awscc_evidently_feature" "example" {
  name        = "example"
  project     = awscc_evidently_project.example.name
  description = "example description"

  variations = [
    {
      variation_name = "Variation1"
      string_value   = "example"
    }
  ]

  tags = [
    {
      key   = "ModifiedBy"
      value = "AWSCC"
    }
  ]
}

resource "awscc_evidently_project" "example" {
  name        = "Example"
  description = "Example Description"

  tags = [
    {
      key   = "ModifiedBy"
      value = "AWSCC"
    }
  ]
}