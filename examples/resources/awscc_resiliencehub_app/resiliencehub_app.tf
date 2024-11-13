resource "awscc_resiliencehub_app" "main" {
  name             = "test-app"
  resource_mappings = [ {
      physical_resource_id = {
        identifier = "s3://terraform-state-files/test-app.tfstate"
        type = "Native"
        # Optional aws_region = ""
        # Optional aws_account_id = ""
      }
      mapping_type = "Terraform"
      # Optional  resource_name = ""
      #           resource_type = ""
      #           logical_resource = { 
      #             logical_resource_name = ""
      #             resource_mapping_name = ""
      #             resource_mapping_arn = awscc_resiliencehub_app.app.arn 
      }]
  resiliency_policy_arn = "arn:aws:resiliencehub:us-west-1:<account-id>:resiliency-policy/<id>"
  app_template_body = <<EOF
    { "resources": [], "appComponents": [], "excludedResources": { "logicalResourceIds": [] }, "version": 2 }
EOF
  description = "This is a test app."
  app_assessment_schedule = "Daily"
  event_subscriptions = []
  permission_model = {
    type = "RoleBased"
    # Optional invoker_role_name = "resilience-hub-test-app-role"
  }
  tags = [{
    key   = "Modified By"
    value = "AWSCC"
  }]
}

