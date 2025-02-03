# Data sources for AWS account and region information
data "aws_caller_identity" "current" {}
data "aws_region" "current" {}

# Create AppConfig application first
resource "awscc_appconfig_application" "example" {
  name        = "example-app-test-123"
  description = "Example AppConfig Application"
}

# Create AppConfig configuration profile
resource "awscc_appconfig_configuration_profile" "example" {
  application_id = awscc_appconfig_application.example.application_id
  name           = "example-profile"
  location_uri   = "hosted"
  description    = "Example configuration profile"
}

# Create the hosted configuration version
resource "awscc_appconfig_hosted_configuration_version" "example" {
  application_id           = awscc_appconfig_application.example.application_id
  configuration_profile_id = awscc_appconfig_configuration_profile.example.configuration_profile_id
  content_type             = "application/json"
  content = jsonencode({
    "environment" : "production",
    "database" : {
      "connection_timeout" : 30,
      "retry_attempts" : 3
    }
  })
  description   = "Example hosted configuration version"
  version_label = "v1.0.0"
}