data "aws_caller_identity" "current" {}

data "aws_region" "current" {}

resource "awscc_codepipeline_custom_action_type" "example" {
  category      = "Build"
  provider_name = "MyCompany"
  version       = "1"

  input_artifact_details = {
    maximum_count = 1
    minimum_count = 1
  }

  output_artifact_details = {
    maximum_count = 1
    minimum_count = 1
  }

  configuration_properties = [
    {
      name        = "BuildCommand"
      description = "Command to run the build"
      type        = "String"
      required    = true
      key         = true
      secret      = false
      queryable   = true
    },
    {
      name        = "BuildEnv"
      description = "Build environment"
      type        = "String"
      required    = true
      key         = false
      secret      = false
      queryable   = false
    }
  ]

  settings = {
    entity_url_template    = "https://build.${data.aws_region.current.name}.${data.aws_caller_identity.current.account_id}.example.com/builds/{Config:BuildCommand}"
    execution_url_template = "https://build.${data.aws_region.current.name}.${data.aws_caller_identity.current.account_id}.example.com/builds/{Config:BuildCommand}/execution"
    revision_url_template  = "https://build.${data.aws_region.current.name}.${data.aws_caller_identity.current.account_id}.example.com/builds/{Config:BuildCommand}/revision"
  }

  tags = [{
    key   = "Modified By"
    value = "AWSCC"
  }]
}