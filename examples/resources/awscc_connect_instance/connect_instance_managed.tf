resource "awscc_connect_instance" "example" {
  identity_management_type = "CONNECT_MANAGED"
  attributes = {
    inbound_calls    = true
    outbound_calls   = true
    contact_lens     = true
    early_media      = true
    contactflow_logs = true
  }
  instance_alias = "example-managed"

  tags = [{
    key   = "Modified By"
    value = "AWSCC"
  }]
}
