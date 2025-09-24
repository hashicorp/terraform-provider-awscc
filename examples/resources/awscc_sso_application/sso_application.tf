resource "awscc_sso_application" "example" {
  name                     = "example"
  application_provider_arn = "arn:aws:sso::aws:applicationProvider/custom"
  instance_arn             = tolist(data.aws_ssoadmin_instances.example.arns)[0]
  description              = "Example application"
  status                   = "ENABLED"
  portal_options = {
    sign_in_options = {
      application_url = "http://www.example.com"
      origin          = "APPLICATION"
    }
    visibility = "ENABLED"

  }
  tags = [{
    key   = "Modified_By"
    value = "AWSCC"
  }]
}

data "aws_ssoadmin_instances" "example" {}
