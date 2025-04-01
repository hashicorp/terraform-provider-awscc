data "aws_caller_identity" "current" {}

# First, set up the IoT audit configuration
resource "awscc_iot_account_audit_configuration" "config" {
  account_id = data.aws_caller_identity.current.account_id
  audit_check_configurations = {
    authenticated_cognito_role_overly_permissive_check = {
      enabled = true
    }
    ca_certificate_expiring_check = {
      enabled = true
    }
    conflicting_client_ids_check = {
      enabled = true
    }
    device_certificate_expiring_check = {
      enabled = true
    }
    iot_policy_overly_permissive_check = {
      enabled = true
    }
  }
  audit_notification_target_configurations = {
    sns = {
      enabled    = true
      role_arn   = awscc_iam_role.audit_notification.arn
      target_arn = awscc_sns_topic.audit_notifications.topic_arn
    }
  }
  role_arn = awscc_iam_role.audit_role.arn
}

# SNS topic for audit notifications
resource "awscc_sns_topic" "audit_notifications" {
  topic_name = "iot-audit-notifications"
  tags = [{
    key   = "Modified By"
    value = "AWSCC"
  }]
}

# IAM role for IoT audit
resource "awscc_iam_role" "audit_role" {
  role_name = "iot-audit-role"
  assume_role_policy_document = jsonencode({
    Version = "2012-10-17"
    Statement = [{
      Action = "sts:AssumeRole"
      Effect = "Allow"
      Principal = {
        Service = "iot.amazonaws.com"
      }
    }]
  })
  tags = [{
    key   = "Modified By"
    value = "AWSCC"
  }]
}

# Policy for IoT audit role
resource "awscc_iam_role_policy" "audit_policy" {
  policy_name = "iot-audit-policy"
  role_name   = awscc_iam_role.audit_role.role_name
  policy_document = jsonencode({
    Version = "2012-10-17"
    Statement = [{
      Effect = "Allow"
      Action = [
        "iot:GetLoggingOptions",
        "iot:GetPolicy",
        "iot:GetPolicy*",
        "iot:ListCertificates",
        "iot:ListPolicies",
        "iot:ListThings",
        "iot:DescribeCertificate",
        "iot:DescribeRoleAlias",
        "iot:GetEffectivePolicies",
        "iot:ListRoleAliases",
        "iot:ListTargetsForPolicy",
        "iot:ListPrincipalThings",
        "iot:ListPolicyVersions",
        "iot:ListThingsInBillingGroup",
        "iot:ListThingsInThingGroup",
        "iot:ListAttachedPolicies",
        "iot:ListCertificatesByCA",
        "iot:ListIndicesForPolicy",
        "iot:ListPrincipalPolicies",
        "iot:ValidateSecurityProfileBehaviors"
      ]
      Resource = "*"
    }]
  })
}

# IAM role for audit notifications
resource "awscc_iam_role" "audit_notification" {
  role_name = "iot-audit-notification-role"
  assume_role_policy_document = jsonencode({
    Version = "2012-10-17"
    Statement = [{
      Action = "sts:AssumeRole"
      Effect = "Allow"
      Principal = {
        Service = "iot.amazonaws.com"
      }
    }]
  })
  tags = [{
    key   = "Modified By"
    value = "AWSCC"
  }]
}

# Create the scheduled audit
resource "awscc_iot_scheduled_audit" "example" {
  depends_on = [awscc_iot_account_audit_configuration.config]
  frequency  = "DAILY"
  target_check_names = [
    "AUTHENTICATED_COGNITO_ROLE_OVERLY_PERMISSIVE_CHECK",
    "CA_CERTIFICATE_EXPIRING_CHECK",
    "CONFLICTING_CLIENT_IDS_CHECK",
    "DEVICE_CERTIFICATE_EXPIRING_CHECK",
    "IOT_POLICY_OVERLY_PERMISSIVE_CHECK"
  ]
  scheduled_audit_name = "example-scheduled-audit"
  tags = [{
    key   = "Modified By"
    value = "AWSCC"
  }]
}