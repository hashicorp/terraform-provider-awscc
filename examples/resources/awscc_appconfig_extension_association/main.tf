# Data sources for AWS account and region
data "aws_caller_identity" "current" {}
# Note: Using data.aws_region.current.region (AWS provider v6.0+)
# For AWS provider < v6.0, use data.aws_region.current.name instead
data "aws_region" "current" {}

# Create AppConfig Application
resource "awscc_appconfig_application" "example" {
  name = "example-app"
  tags = [{
    key   = "Modified By"
    value = "AWSCC"
  }]
}

# Create AppConfig Environment
resource "awscc_appconfig_environment" "example" {
  name           = "example-env"
  application_id = awscc_appconfig_application.example.id
  description    = "Example environment"
  tags = [{
    key   = "Modified By"
    value = "AWSCC"
  }]
}

# Create AppConfig Extension Association using AWS Lambda pre-built extension
resource "awscc_appconfig_extension_association" "example" {
  extension_identifier = "arn:aws:appconfig:${data.aws_region.current.region}:aws:lambda:1"
  resource_identifier = "arn:aws:appconfig:${data.aws_region.current.region}:${data.aws_caller_identity.current.account_id}:application/${awscc_appconfig_application.example.id}/environment/${awscc_appconfig_environment.example.id}"
  parameters = {
    "FunctionARN" = "arn:aws:lambda:${data.aws_region.current.region}:${data.aws_caller_identity.current.account_id}:function:example-function"
  }
  tags = [{
    key   = "Modified By"
    value = "AWSCC"
  }]
}