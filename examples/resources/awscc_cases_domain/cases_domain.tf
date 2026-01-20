resource "awscc_cases_domain" "example" {
  name = "example-cases-domain"

  tags = [{
    key   = "Environment"
    value = "test"
  }]
}
