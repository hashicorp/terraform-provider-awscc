resource "awscc_connect_instance" "example" {
  identity_management_type = "CONNECT_MANAGED"
  attributes = {
    inbound_calls  = true
    outbound_calls = true
  }
  instance_alias = "example-connect-instance"

  tags = [{
    key   = "Environment"
    value = "Example"
    }, {
    key   = "Name"
    value = "example-connect-instance"
  }]
}

resource "awscc_connect_contact_flow_module" "example" {
  instance_arn = awscc_connect_instance.example.arn
  name         = "example-contact-flow-module"
  description  = "Example Contact Flow Module for testing alias"

  content = jsonencode({
    Version     = "2019-10-30"
    StartAction = "12345678-1234-1234-1234-123456789012"
    Actions = [
      {
        Identifier = "12345678-1234-1234-1234-123456789012"

        Parameters = {
          Text = "Hello from example contact flow module"
        }

        Transitions = {
          NextAction = "abcdef-abcd-abcd-abcd-abcdefghijkl"
          Errors     = []
          Conditions = []
        }

        Type = "MessageParticipant"
      },
      {
        Identifier  = "abcdef-abcd-abcd-abcd-abcdefghijkl"
        Type        = "DisconnectParticipant"
        Parameters  = {}
        Transitions = {}
      }
    ]
    Settings = {
      InputParameters  = []
      OutputParameters = []
      Transitions = [
        {
          DisplayName   = "Success"
          ReferenceName = "Success"
          Description   = ""
        },
        {
          DisplayName   = "Error"
          ReferenceName = "Error"
          Description   = ""
        }
      ]
    }
  })

  tags = [{
    key   = "Environment"
    value = "Example"
    }, {
    key   = "Name"
    value = "example-contact-flow-module"
  }]
}

resource "awscc_connect_contact_flow_module_version" "example" {
  contact_flow_module_id = awscc_connect_contact_flow_module.example.contact_flow_module_arn
  description            = "Example contact flow module version"
}

resource "awscc_connect_contact_flow_module_alias" "example" {
  contact_flow_module_id      = awscc_connect_contact_flow_module.example.contact_flow_module_arn
  contact_flow_module_version = awscc_connect_contact_flow_module_version.example.version
  name                        = "example-alias"
  description                 = "Example contact flow module alias"
}
