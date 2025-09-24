# Create an Amazon Connect instance first
resource "awscc_connect_instance" "example" {
  attributes = {
    inbound_calls  = true
    outbound_calls = true
  }
  identity_management_type = "CONNECT_MANAGED"
  instance_alias           = "example-security-profile-instance"
  tags = [{
    key   = "Modified By"
    value = "AWSCC"
  }]
}

# Create the security profile
resource "awscc_connect_security_profile" "example" {
  instance_arn          = awscc_connect_instance.example.arn
  security_profile_name = "example-security-profile"
  description           = "Example security profile created via AWSCC"

  permissions = [
    "BasicAgentAccess",
    "OutboundCallAccess"
  ]

  applications = [{
    namespace               = "Connect"
    application_permissions = ["AccessAgentApplication"]
  }]

  tags = [{
    key   = "Modified By"
    value = "AWSCC"
  }]
}