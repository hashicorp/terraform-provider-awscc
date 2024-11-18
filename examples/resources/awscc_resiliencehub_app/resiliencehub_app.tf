resource "awscc_resiliencehub_app" "main" {
  name        = "test-app"
  description = "This is a test app."
  app_template_body = jsonencode({
    format_version = 2
    resources = [{
      name = "test-app"
      type = "AWS::EC2::Instance"
      parameters = {
        tags = {
          env = "dev"
        }
      }
    }]
  })
  resiliency_policy_arn = awscc_resiliencehub_resiliency_policy.test.arn
  resource_mappings = [{
    physical_resource_id = {
      identifier     = "s3://terraform-state-files/test-app.tfstate"
      type           = "Native"
      aws_region     = "us-west-2"
      aws_account_id = "112223333444"
    }
    mapping_type  = "Terraform"
    resource_name = "test-app"
    resource_type = "AWS::EC2::Instance"
    logical_resource = {
      logical_resource_name = "test-app"
      resource_mapping_arn  = awscc_resiliencehub_app.app.arn
    }

  }]
  event_subscriptions = [{
    name          = "test-app"
    sns_topic_arn = awscc_sns_topic.sns_example.arn
  }]
  app_assessment_schedule = "Daily"
  permission_model = {
    type              = "RoleBased"
    invoker_role_name = "resilience-hub-test-app-role"
  }
  tags = [{
    key   = "Modified By"
    value = "AWSCC"
  }]
}
