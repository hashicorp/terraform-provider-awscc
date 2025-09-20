resource "awscc_smsvoice_sender_id" "example" {
  sender_id        = "MYCOMPANY"
  iso_country_code = "GB"

  deletion_protection_enabled = false

  tags = [
    {
      key   = "Environment"
      value = "example"
    },
    {
      key   = "Name"
      value = "example-sender-id"
    }
  ]
}