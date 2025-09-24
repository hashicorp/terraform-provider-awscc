resource "awscc_glue_registry" "example" {
  name        = "example-registry"
  description = "Glue registry example"

  tags = [{
    key   = "Modified By"
    value = "AWSCC"
  }]
}