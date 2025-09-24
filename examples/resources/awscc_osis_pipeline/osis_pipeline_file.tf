resource "awscc_osis_pipeline" "example_pipeline" {
  pipeline_name = "example-pipeline"
  min_units     = 1
  max_units     = 4

  pipeline_configuration_body = file("example-pipeline.yaml")
}