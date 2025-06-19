# Data sources for dynamic values
data "aws_caller_identity" "current" {}
data "aws_region" "current" {}

# S3 bucket for Guard rules and validation reports
resource "awscc_s3_bucket" "guard_bucket" {
  bucket_name = "cloudformation-guard-rules-${random_id.bucket_suffix.hex}"

  tags = [{
    key   = "Modified By"
    value = "AWSCC"
  }]
}

resource "random_id" "bucket_suffix" {
  byte_length = 4
}

resource "awscc_iam_role" "guard_hook_role" {
  role_name = "CloudFormationGuardHookRole"

  assume_role_policy_document = jsonencode({
    Version = "2012-10-17"
    Statement = [
      {
        Effect = "Allow"
        Principal = {
          Service = "hooks.cloudformation.amazonaws.com"
        }
        Action = "sts:AssumeRole"
      }
    ]
  })

  policies = [{
    policy_name = "GuardHookExecutionPolicy"
    policy_document = jsonencode({
      Version = "2012-10-17"
      Statement = [
        {
          Effect = "Allow"
          Action = [
            "s3:GetObject",
            "s3:GetObjectVersion"
          ]
          Resource = "${awscc_s3_bucket.guard_bucket.arn}/*"
        },
        {
          Effect = "Allow"
          Action = [
            "s3:PutObject",
            "s3:PutObjectAcl"
          ]
          Resource = "${awscc_s3_bucket.guard_bucket.arn}/reports/*"
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
  }]

  tags = [{
    key   = "Modified By"
    value = "AWSCC"
  }]
}

# Upload sample Guard rule file
resource "aws_s3_object" "sample_guard_rule" {
  bucket = awscc_s3_bucket.guard_bucket.bucket_name
  key    = "rules/security-rules.guard"

  content = <<-EOT
    # Sample Guard rules for security validation
    
    # Ensure Lambda functions use supported runtimes
    AWS::Lambda::Function {
        Properties.Runtime in [
            "python3.9", "python3.10", "python3.11", "python3.12",
            "nodejs18.x", "nodejs20.x",
            "java11", "java17", "java21",
            "dotnet6", "dotnet8"
        ]
    }
  EOT

  content_type = "text/plain"

  tags = {
    "Modified By" = "AWSCC"
  }
}

# CloudFormation Guard Hook
resource "awscc_cloudformation_guard_hook" "security_validation_hook" {
  alias          = "CCAPI::Lambda::Hooks"
  execution_role = awscc_iam_role.guard_hook_role.arn

  rule_location = {
    uri = "s3://${awscc_s3_bucket.guard_bucket.bucket_name}/rules/"
  }

  target_operations = ["CLOUD_CONTROL"]

  failure_mode = "FAIL"
  hook_status  = "ENABLED"
  log_bucket   = awscc_s3_bucket.guard_bucket.bucket_name

  # Target specific resource types for validation
  target_filters = {
    actions           = ["CREATE", "UPDATE"]
    invocation_points = ["PRE_PROVISION"]
    target_names      = ["AWS::Lambda::Function"]
  }
}

