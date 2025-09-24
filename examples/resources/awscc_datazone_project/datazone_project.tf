resource "awscc_datazone_project" "example" {
  domain_identifier = awscc_datazone_domain.example.id
  name              = "example"
}

resource "awscc_datazone_domain" "example" {
  name                  = "example"
  domain_execution_role = awscc_iam_role.awscc_datazone_role.arn
  description           = "Datazone domain example"

  tags = [{
    key   = "Modified By"
    value = "AWSCC"
  }]
}