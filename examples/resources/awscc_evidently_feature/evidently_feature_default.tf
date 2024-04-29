resource "awscc_evidently_feature" "example" {
  name              = "example"
  project           = awscc_evidently_project.example.name
  default_variation = "Variation2"

  variations = [
    {
      variation_name = "Variation1"
      string_value   = "exampleval1"
    },
    {

      variation_name = "Variation2"
      string_value   = "exampleval2"
    }
  ]

  tags = [
    {
      key   = "ModifiedBy"
      value = "AWSCC"
    }
  ]


}