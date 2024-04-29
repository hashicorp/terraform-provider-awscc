resource "awscc_evidently_feature" "example" {
  name                = "example"
  project             = awscc_evidently_project.example.name
  evaluation_strategy = "ALL_RULES"
  entity_overrides = [
    {
      entity_id = "test1"
      variation = "Variation1"
    }
  ]
  variations = [
    {
      variation_name = "Variation1"
      string_value   = "exampleval1"
    }
  ]

  tags = [
    {
      key   = "ModifiedBy"
      value = "AWSCC"
    }
  ]

}