resource "awscc_smsvoice_protect_configuration" "example" {
  tags = [
    {
      key   = "Name"
      value = "example-protect-config"
    },
    {
      key   = "Environment"
      value = "example"
    }
  ]
}