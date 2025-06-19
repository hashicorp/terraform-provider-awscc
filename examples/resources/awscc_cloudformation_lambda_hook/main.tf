data "aws_caller_identity" "current" {}
data "aws_region" "current" {}

data "archive_file" "lambda_zip" {
  type        = "zip"
  output_path = "/tmp/lambda_hook_function.zip"
  source {
    content  = <<EOF
import json
import logging

logger = logging.getLogger()
logger.setLevel(logging.INFO)

def lambda_handler(event, context):
    """
    CloudFormation Lambda Hook handler
    Processes CloudFormation stack operation events
    """
    logger.info(f"Received CloudFormation hook event: {json.dumps(event)}")
    
    # Extract hook details
    hook_type = event.get('HookType', 'Unknown')
    stack_name = event.get('StackName', 'Unknown')
    operation = event.get('Operation', 'Unknown')
    
    logger.info(f"Processing {hook_type} hook for stack: {stack_name}, operation: {operation}")
    
    # Hook logic - validate or monitor the operation
    # For this example, we'll just log and allow all operations
    
    response = {
        'Status': 'SUCCESS',
        'Message': f'Hook processed successfully for stack {stack_name}'
    }
    
    logger.info(f"Hook response: {json.dumps(response)}")
    return response
EOF
    filename = "lambda_function.py"
  }
}

resource "awscc_iam_role" "lambda_role" {
  role_name = "cloudformation-lambda-hook-function-role"
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

  managed_policy_arns = [
    "arn:aws:iam::aws:policy/service-role/AWSLambdaBasicExecutionRole"
  ]

  tags = [{
    key   = "Modified By"
    value = "AWSCC"
  }]
}

# Lambda function for the hook
resource "aws_lambda_function" "hook_function" {
  filename      = data.archive_file.lambda_zip.output_path
  function_name = "cloudformation-lambda-hook-processor"
  role          = awscc_iam_role.lambda_role.arn
  handler       = "lambda_function.lambda_handler"
  runtime       = "python3.11"
  timeout       = 60

  source_code_hash = data.archive_file.lambda_zip.output_base64sha256

}

# IAM role for CloudFormation Lambda Hook execution
resource "awscc_iam_role" "hook_execution_role" {
  role_name = "cloudformation-lambda-hook-execution-role"
  assume_role_policy_document = jsonencode({
    Version = "2012-10-17"
    Statement = [
      {
        Action = "sts:AssumeRole"
        Effect = "Allow"
        Principal = {
          Service = "hooks.cloudformation.amazonaws.com"
        }
      }
    ]
  })

  tags = [{
    key   = "Modified By"
    value = "AWSCC"
  }]
}

# IAM policy for hook execution role
resource "awscc_iam_role_policy" "hook_execution_policy" {
  role_name   = awscc_iam_role.hook_execution_role.role_name
  policy_name = "CloudFormationLambdaHookExecutionPolicy"
  policy_document = jsonencode({
    Version = "2012-10-17"
    Statement = [
      {
        Effect = "Allow"
        Action = [
          "lambda:InvokeFunction"
        ]
        Resource = aws_lambda_function.hook_function.arn
      },
      {
        Effect = "Allow"
        Action = [
          "logs:CreateLogGroup",
          "logs:CreateLogStream",
          "logs:PutLogEvents"
        ]
        Resource = "arn:aws:logs:${data.aws_region.current.name}:${data.aws_caller_identity.current.account_id}:*"
      }
    ]
  })
}

# CloudFormation Lambda Hook
resource "awscc_cloudformation_lambda_hook" "example" {
  alias           = "CCAPI::TEST::Hooks"
  execution_role  = awscc_iam_role.hook_execution_role.arn
  failure_mode    = "WARN"
  lambda_function = aws_lambda_function.hook_function.arn
  hook_status     = "ENABLED"

  target_operations = ["CLOUD_CONTROL"]

  # Target specific resource types and operations
  target_filters = {
    actions = [
      "CREATE",
      "UPDATE",
      "DELETE"
    ]
    invocation_points = [
      "PRE_PROVISION"
    ]
    target_names = [
      "AWS::S3::Bucket",
      "AWS::EC2::Instance",
      "AWS::Lambda::Function"
    ]
  }

}
