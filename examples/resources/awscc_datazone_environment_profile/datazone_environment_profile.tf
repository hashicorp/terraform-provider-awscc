resource "awscc_datazone_environment_profile" "example" {
  name                             = "example"
  description                      = "Example environment profile"
  aws_account_id                   = data.aws_caller_identity.current.account_id
  aws_account_region               = "us-east-1"
  domain_identifier                = awscc_datazone_domain.example.domain_id
  environment_blueprint_identifier = awscc_datazone_environment_blueprint_configuration.example.environment_blueprint_id
  project_identifier               = awscc_datazone_project.example.project_id
}

data "aws_caller_identity" "current" {}