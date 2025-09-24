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