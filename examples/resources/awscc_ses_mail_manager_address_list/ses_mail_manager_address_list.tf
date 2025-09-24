# SES Mail Manager address list for email filtering and management
resource "awscc_ses_mail_manager_address_list" "example" {
  address_list_name = "example-address-list" # Name for the address list

  tags = [
    {
      key   = "Environment"
      value = "example" # Environment designation
    },
    {
      key   = "Name"
      value = "example-address-list" # Resource identifier
    }
  ]
}
