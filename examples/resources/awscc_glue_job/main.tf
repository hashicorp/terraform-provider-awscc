# Data sources for AWS account ID and region
data "aws_caller_identity" "current" {}
data "aws_region" "current" {}

# IAM role policy document
data "aws_iam_policy_document" "trust" {
  statement {
    effect = "Allow"
    actions = [
      "sts:AssumeRole"
    ]
    principals {
      type = "Service"
      identifiers = [
        "glue.amazonaws.com"
      ]
    }
  }
}

# IAM role policy for Glue job
data "aws_iam_policy_document" "glue_job" {
  statement {
    effect = "Allow"
    actions = [
      "s3:GetObject",
      "s3:PutObject",
      "logs:CreateLogGroup",
      "logs:CreateLogStream",
      "logs:PutLogEvents"
    ]
    resources = [
      "arn:aws:s3:::aws-glue-*/*",
      "arn:aws:logs:${data.aws_region.current.name}:${data.aws_caller_identity.current.account_id}:log-group:/aws-glue/*"
    ]
  }
}

# IAM role for Glue job
resource "awscc_iam_role" "glue_job" {
  assume_role_policy_document = data.aws_iam_policy_document.trust.json
  path                        = "/service-role/"
  policies = [
    {
      policy_document = data.aws_iam_policy_document.glue_job.json
      policy_name     = "GlueJobPolicy"
    }
  ]
  role_name = "ExampleGlueJobRole"
  tags = [{
    key   = "Modified By"
    value = "AWSCC"
  }]
}

# Glue job
resource "awscc_glue_job" "example" {
  name        = "example-glue-job"
  description = "Example Glue Job created with AWSCC provider"
  role        = awscc_iam_role.glue_job.arn

  command = {
    name            = "pythonshell"
    python_version  = "3.9"
    script_location = "s3://aws-glue-scripts-${data.aws_caller_identity.current.account_id}-${data.aws_region.current.name}/example-script.py"
  }

  glue_version = "3.0"
  max_capacity = 0.0625

  execution_property = {
    max_concurrent_runs = 1
  }

  tags = [{
    key   = "Modified By"
    value = "AWSCC"
  }]
}