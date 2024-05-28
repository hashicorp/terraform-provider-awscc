resource "awscc_appconfig_environment" "example" {
  application_id = awscc_appconfig_application.example.application_id
  name           = "example"
  description    = "Example environment"


  tags = [{
    key   = "ModifiedBy"
    value = "AWSCC"
  }]
}