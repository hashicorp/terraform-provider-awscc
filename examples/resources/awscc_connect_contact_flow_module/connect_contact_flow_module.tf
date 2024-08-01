resource "awscc_connect_contact_flow_module" "example" {
  instance_arn = awscc_connect_instance.example.arn
  name         = "Example"
  description  = "Example Contact Flow Module Description"

  content = jsonencode({
    Version     = "2019-10-30"
    StartAction = "12345678-1234-1234-1234-123456789012"
    Actions = [
      {
        Identifier = "12345678-1234-1234-1234-123456789012"

        Parameters = {
          Text = "Hello contact flow module"
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
    key   = "ModifiedBy"
    value = "AWSCC"
  }]
}
