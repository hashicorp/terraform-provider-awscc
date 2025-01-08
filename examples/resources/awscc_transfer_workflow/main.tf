# Get current AWS account ID and region
data "aws_caller_identity" "current" {}
data "aws_region" "current" {}

# Create an S3 bucket for workflow operations
resource "awscc_s3_bucket" "workflow_bucket" {
  bucket_name = "transfer-workflow-example-${data.aws_caller_identity.current.account_id}"
}

# Create IAM role for the Lambda function
resource "awscc_iam_role" "workflow_lambda_role" {
  role_name = "transfer-workflow-lambda-role"
  assume_role_policy_document = jsonencode({
    Version = "2012-10-17"
    Statement = [
      {
        Action = "sts:AssumeRole"
        Effect = "Allow"
        Principal = {
          Service = "lambda.amazonaws.com"
        }
      }
    ]
  })

  policies = [
    {
      policy_name = "lambda-execution-policy"
      policy_document = jsonencode({
        Version = "2012-10-17"
        Statement = [
          {
            Effect = "Allow"
            Action = [
              "logs:CreateLogGroup",
              "logs:CreateLogStream",
              "logs:PutLogEvents"
            ]
            Resource = "arn:aws:logs:${data.aws_region.current.name}:${data.aws_caller_identity.current.account_id}:log-group:/aws/lambda/*"
          }
        ]
      })
    }
  ]
}

# Create Lambda function for custom step
resource "aws_lambda_function" "workflow_processor" {
  filename      = "lambda.zip"
  function_name = "transfer-workflow-processor"
  role          = awscc_iam_role.workflow_lambda_role.arn
  handler       = "index.handler"
  runtime       = "nodejs16.x"

  lifecycle {
    ignore_changes = [filename]
  }
}

# Create the Transfer Workflow
resource "awscc_transfer_workflow" "example" {
  description = "Example Transfer Workflow"

  steps = [
    {
      type = "COPY"
      copy_step_details = {
        name                 = "copy-to-processing"
        source_file_location = "$${original.file}"
        destination_file_location = {
          s3_file_location = {
            bucket = awscc_s3_bucket.workflow_bucket.id
            key    = "processing/$${original.file}"
          }
        }
      }
    },
    {
      type = "CUSTOM"
      custom_step_details = {
        name                 = "process-file"
        target               = aws_lambda_function.workflow_processor.arn
        timeout_seconds      = 60
        source_file_location = "$${original.file}"
      }
    }
  ]

  on_exception_steps = [
    {
      type = "COPY"
      copy_step_details = {
        name                 = "copy-to-error"
        source_file_location = "$${original.file}"
        destination_file_location = {
          s3_file_location = {
            bucket = awscc_s3_bucket.workflow_bucket.id
            key    = "error/$${original.file}"
          }
        }
      }
    }
  ]

  tags = [{
    key   = "Environment"
    value = "Example"
  }]
}