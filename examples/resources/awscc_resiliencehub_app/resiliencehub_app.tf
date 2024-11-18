resource "awscc_resiliencehub_app" "example" {
  name        = "test-app"
  description = "This is a test app."
  app_template_body = jsonencode({
    format_version = 2
    resources = [{
      name = ""
      type = ""
      parameters = {
        tags = {
          env = "dev"
        }
      }
    }]
  })
  resiliency_policy_arn = awscc_resiliencehub_resiliency_policy.example.arn
  resource_mappings = [{
    physical_resource_id = {
      identifier     = "s3://terraform-state-files/test-app.tfstate"
      type           = "Native"
      aws_region     = ""
      aws_account_id = ""
    }
    mapping_type  = "Terraform"
    resource_name = ""
    resource_type = ""
    logical_resource = {
      logical_resource_name = ""
      resource_mapping_name = ""
      resource_mapping_arn  = awscc_resiliencehub_app.example.arn
    }
  }]
  app_assessment_schedule = "Daily"
  event_subscriptions     = []
  permission_model = {
    type              = "RoleBased"
    invoker_role_name = "resilience-hub-test-app-role"
  }
  tags = [{
    key   = "Modified By"
    value = "AWSCC"
  }]
}
