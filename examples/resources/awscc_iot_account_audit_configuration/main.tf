# Get current account ID
data "aws_caller_identity" "current" {}

# Create IAM role for IoT Device Defender using AWSCC provider
resource "awscc_iam_role" "iot_audit_role" {
  role_name = "IoTDeviceDefenderAuditRole"
  assume_role_policy_document = jsonencode({
    Version = "2012-10-17"
    Statement = [
      {
        Action = "sts:AssumeRole"
        Effect = "Allow"
        Principal = {
          Service = "iot.amazonaws.com"
        }
      }
    ]
  })

  tags = [{
    key   = "Modified By"
    value = "AWSCC"
  }]
}

resource "awscc_iam_role_policy" "iot_audit_policy" {
  policy_name = "IoTDeviceDefenderAuditAccess"
  role_name   = awscc_iam_role.iot_audit_role.role_name

  policy_document = jsonencode({
    Version = "2012-10-17"
    Statement = [
      {
        Effect = "Allow"
        Action = [
          "iot:GetLoggingOptions",
          "iot:GetPolicy",
          "iot:GetCACertificate",
          "iot:ListCACertificates",
          "iot:ListCertificates",
          "iot:GetCertificate",
          "iot:ListRoleAliases",
          "iot:DescribeRoleAlias",
          "iam:GetRole"
        ]
        Resource = "*"
      }
    ]
  })
}

# Create IoT Device Defender audit configuration
resource "awscc_iot_account_audit_configuration" "example" {
  account_id = data.aws_caller_identity.current.account_id
  role_arn   = awscc_iam_role.iot_audit_role.arn

  audit_check_configurations = {
    authenticated_cognito_role_overly_permissive_check = {
      enabled = true
    }
    ca_certificate_expiring_check = {
      enabled = true
    }
    device_certificate_expiring_check = {
      enabled = true
    }
    iot_policy_overly_permissive_check = {
      enabled = true
    }
    logging_disabled_check = {
      enabled = true
    }
  }

  audit_notification_target_configurations = {
    sns = {
      enabled    = false
      role_arn   = null
      target_arn = null
    }
  }
}