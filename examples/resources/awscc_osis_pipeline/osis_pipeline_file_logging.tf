resource "awscc_logs_log_group" "log_group" {
  log_group_name = "/aws/vendedlogs/OpenSearchIngestion/example-pipeline/audit-logs"
}

resource "awscc_osis_pipeline" "example_pipeline" {
  pipeline_name = "example-pipeline"
  min_units     = 1
  max_units     = 4

  pipeline_configuration_body = file("example-pipeline.yaml")

  log_publishing_options = {
    is_logging_enabled = true
    cloudwatch_log_destination = {
      log_group = awscc_logs_log_group.log_group.log_group_name
    }
  }
}