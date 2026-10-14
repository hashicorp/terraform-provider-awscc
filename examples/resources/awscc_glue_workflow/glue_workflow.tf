data "aws_region" "current" {}

resource "awscc_glue_workflow" "example" {
  name                = "example-glue-workflow"
  description         = "Example Glue workflow orchestrating ETL jobs and crawlers"
  max_concurrent_runs = 1

  default_run_properties = jsonencode({
    environment = "dev"
    region      = data.aws_region.current.region
  })

  tags = jsonencode({
    "Modified By" = "AWSCC"
  })
}

output "workflow_name" {
  description = "The name of the Glue workflow"
  value       = awscc_glue_workflow.example.id
}

output "workflow_description" {
  description = "The description of the Glue workflow"
  value       = awscc_glue_workflow.example.description
}
