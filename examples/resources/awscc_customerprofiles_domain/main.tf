# Customer Profiles Domain
resource "awscc_customerprofiles_domain" "example" {
  domain_name             = "example-domain-awscc-test"
  default_expiration_days = 365
  tags = [{
    key   = "ModifiedBy"
    value = "AWSCC"
  }]
}