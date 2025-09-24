# Example Organizations Account
resource "awscc_organizations_account" "example" {
  account_name = "example-member-account"
  email        = "example-account@example.com"
  role_name    = "OrganizationAccountAccessRole"

  tags = [{
    key   = "Modified By"
    value = "AWSCC"
  }]
}