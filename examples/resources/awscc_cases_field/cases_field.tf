resource "awscc_cases_field" "example" {
  domain_id = awscc_cases_domain.example.domain_id
  name      = "example-field"
  type      = "Text"

  tags = [{
    key   = "Environment"
    value = "test"
  }]
}

resource "awscc_cases_domain" "example" {
  name = "example-cases-domain-for-field"

  tags = [{
    key   = "Environment"
    value = "test"
  }]
}
