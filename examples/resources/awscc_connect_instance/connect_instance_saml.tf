resource "awscc_connect_instance" "example" {
  identity_management_type = "SAML"
  attributes = {
    inbound_calls  = true
    outbound_calls = true
  }
  instance_alias = "example-saml"

  tags = [{
    key   = "Modified By"
    value = "AWSCC"
  }]
}
