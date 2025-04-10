# Create a Cognito User Pool (using AWS provider as AWSCC version not available)
resource "aws_cognito_user_pool" "example" {
  name = "example-user-pool"
  # Note: The Cognito User Pool must be in ENHANCED tier for log delivery to work
  user_pool_add_ons {
    advanced_security_mode = "OFF"
  }

  password_policy {
    minimum_length    = 8
    require_lowercase = true
    require_numbers   = true
    require_symbols   = true
    require_uppercase = true
  }

  schema {
    attribute_data_type      = "String"
    developer_only_attribute = false
    mutable                  = true
    name                     = "email"
    required                 = true

    string_attribute_constraints {
      max_length = "2048"
      min_length = "0"
    }
  }

  admin_create_user_config {
    allow_admin_create_user_only = false
  }

  tags = {
    "Modified By" = "AWSCC"
  }
}

# Create CloudWatch Log Group using AWSCC provider
resource "awscc_logs_log_group" "example" {
  log_group_name = "/aws/cognito/example-logs"
  tags = [{
    key   = "Modified By"
    value = "AWSCC"
  }]
}

# IAM role for CloudWatch Logs using AWSCC provider
resource "awscc_iam_role" "cognito_cloudwatch" {
  role_name = "cognito-cloudwatch-role"
  assume_role_policy_document = jsonencode({
    Version = "2012-10-17"
    Statement = [
      {
        Action = "sts:AssumeRole"
        Effect = "Allow"
        Principal = {
          Service = "cognito-idp.amazonaws.com"
        }
      }
    ]
  })
  tags = [{
    key   = "Modified By"
    value = "AWSCC"
  }]
}

# CloudWatch Logs policy using AWSCC provider
resource "awscc_iam_role_policy" "cognito_cloudwatch" {
  policy_name = "cognito-cloudwatch-policy"
  role_name   = awscc_iam_role.cognito_cloudwatch.role_name
  policy_document = jsonencode({
    Version = "2012-10-17"
    Statement = [
      {
        Effect = "Allow"
        Action = [
          "logs:CreateLogStream",
          "logs:PutLogEvents"
        ]
        Resource = awscc_logs_log_group.example.arn
      }
    ]
  })
}

# Log Delivery Configuration
resource "awscc_cognito_log_delivery_configuration" "example" {
  user_pool_id = aws_cognito_user_pool.example.id
  log_configurations = [
    {
      cloudwatch_logs_configuration = {
        log_group_arn = awscc_logs_log_group.example.arn
      }
      event_source = "userAuthEvents"
      log_level    = "INFO"
    }
  ]
}