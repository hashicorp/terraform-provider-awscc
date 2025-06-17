# Get AWS account ID and region
data "aws_caller_identity" "current" {}
data "aws_region" "current" {}

# Create S3 bucket for workflow files
resource "awscc_s3_bucket" "workflow" {
  bucket_name = "omics-workflow-${data.aws_caller_identity.current.account_id}-${data.aws_region.current.name}"
  tags = [{
    key   = "Modified By"
    value = "AWSCC"
  }]
}

# Upload workflow file (keeping AWS standard as no AWSCC equivalent)
resource "aws_s3_object" "workflow" {
  bucket = awscc_s3_bucket.workflow.id
  key    = "workflow.wdl"
  source = "workflow.wdl"
  etag   = filemd5("workflow.wdl")
}

# Create the Omics Workflow
resource "awscc_omics_workflow" "example" {
  name           = "example-workflow"
  description    = "Example Omics Workflow"
  engine         = "WDL"
  definition_uri = "s3://${awscc_s3_bucket.workflow.id}/${aws_s3_object.workflow.key}"
  tags = {
    "Modified By" = "AWSCC"
  }
}

# Create a workflow version
resource "awscc_omics_workflow_version" "example" {
  workflow_id    = awscc_omics_workflow.example.id
  version_name   = "v1.0.0"
  description    = "Example workflow version"
  engine         = "WDL"
  definition_uri = "s3://${awscc_s3_bucket.workflow.id}/${aws_s3_object.workflow.key}"
  parameter_template = {
    "name" = {
      description = "Name to say hello to"
      optional    = true
    }
  }
  tags = {
    "Modified By" = "AWSCC"
  }
}