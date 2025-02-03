# Get current AWS account ID and region
data "aws_caller_identity" "current" {}
data "aws_region" "current" {}

# Create an IAM role for RUM guest access
resource "awscc_iam_role" "rum_guest_role" {
  role_name = "rum-guest-role"
  assume_role_policy_document = jsonencode({
    Version = "2012-10-17"
    Statement = [
      {
        Effect = "Allow"
        Principal = {
          "Federated" : "cognito-identity.amazonaws.com"
        }
        Action = "sts:AssumeRole"
      }
    ]
  })
  policies = [
    {
      policy_document = jsonencode({
        Version = "2012-10-17"
        Statement = [
          {
            Effect = "Allow"
            Action = [
              "rum:PutRumEvents"
            ]
            Resource = "arn:aws:rum:${data.aws_region.current.name}:${data.aws_caller_identity.current.account_id}:appmonitor/example-app-monitor"
          }
        ]
      })
      policy_name = "rum-guest-policy"
    }
  ]
}

# Create Cognito Identity Pool
resource "awscc_cognito_identity_pool" "rum_identity_pool" {
  identity_pool_name               = "rum-identity-pool"
  allow_unauthenticated_identities = true
}

# Create the RUM App Monitor
resource "awscc_rum_app_monitor" "example" {
  name   = "example-app-monitor"
  domain = "example.com"

  app_monitor_configuration = {
    allow_cookies       = true
    enable_x_ray        = false
    session_sample_rate = 0.1
    telemetries         = ["performance", "errors", "http"]
    guest_role_arn      = awscc_iam_role.rum_guest_role.arn
    identity_pool_id    = awscc_cognito_identity_pool.rum_identity_pool.id
    included_pages      = ["https://example.com/*"]
    metric_destinations = [{
      destination = "CloudWatch"
      metric_definitions = [{
        name       = "JsErrorCount"
        namespace  = "RUM/CustomMetrics"
        unit_label = "Count"
        event_pattern = jsonencode({
          event_type = ["com.amazon.rum.js_error_event"]
          metadata = {
            browserName = ["Chrome", "Firefox", "Safari"]
          }
        })
      }]
    }]
  }

  cw_log_enabled = true

  custom_events = {
    status = "ENABLED"
  }

  tags = [{
    key   = "ModifiedBy"
    value = "AWSCC"
  }]
}