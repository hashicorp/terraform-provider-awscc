# Application resource
resource "awscc_appconfig_application" "example" {
  name = "example-app"

  tags = [
    {
      key   = "Environment"
      value = "example"
    },
    {
      key   = "Name"
      value = "example-application"
    }
  ]
}

# Environment resource
resource "awscc_appconfig_environment" "example" {
  name           = "example-env"
  application_id = awscc_appconfig_application.example.application_id

  tags = [
    {
      key   = "Environment"
      value = "example"
    },
    {
      key   = "Name"
      value = "example-environment"
    }
  ]
}

# Configuration Profile resource
resource "awscc_appconfig_configuration_profile" "example" {
  application_id = awscc_appconfig_application.example.application_id
  name           = "example-profile"
  location_uri   = "hosted"

  tags = [
    {
      key   = "Environment"
      value = "example"
    },
    {
      key   = "Name"
      value = "example-configuration-profile"
    }
  ]
}

# Hosted Configuration Version resource
resource "awscc_appconfig_hosted_configuration_version" "example" {
  application_id           = awscc_appconfig_application.example.application_id
  configuration_profile_id = awscc_appconfig_configuration_profile.example.configuration_profile_id
  content_type             = "application/json"

  content = base64encode(jsonencode({
    flags = {
      example_flag = {
        name    = "Example Feature Flag"
        enabled = true
      }
    }
  }))
}

# Deployment resource
resource "awscc_appconfig_deployment" "example" {
  application_id           = awscc_appconfig_application.example.application_id
  environment_id           = awscc_appconfig_environment.example.environment_id
  deployment_strategy_id   = "AppConfig.AllAtOnce"
  configuration_profile_id = awscc_appconfig_configuration_profile.example.configuration_profile_id
  configuration_version    = awscc_appconfig_hosted_configuration_version.example.version_number

  description = "Example deployment of configuration"

  tags = [
    {
      key   = "Environment"
      value = "example"
    },
    {
      key   = "Name"
      value = "example-deployment"
    }
  ]
}
