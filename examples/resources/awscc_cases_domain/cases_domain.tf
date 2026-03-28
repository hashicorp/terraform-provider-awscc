resource "awscc_cases_domain" "example" {
  name = "example-cases-domain"

  tags = [
    {
      key   = "Name"
      value = "example-cases-domain"
    },
    {
      key   = "Environment"
      value = "example"
    }
  ]
}
