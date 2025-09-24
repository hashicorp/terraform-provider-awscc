resource "awscc_connect_contact_flow" "example" {
  instance_arn = awscc_connect_instance.example.arn
  name         = "quick-connect-example"
  type         = "CONTACT_FLOW"
  description  = "example"
  state        = "ACTIVE"
  content = jsonencode({
    Version     = "2019-10-30"
    StartAction = "12345678-1234-1234-1234-123456789012"
    Actions = [
      {
        Identifier = "12345678-1234-1234-1234-123456789012"
        Type       = "MessageParticipant"

        Transitions = {
          NextAction = "abcdef-abcd-abcd-abcd-abcdefghijkl"
          Errors     = []
          Conditions = []
        }

        Parameters = {
          Text = "Thanks for calling the sample flow!"
        }
      },
      {
        Identifier  = "abcdef-abcd-abcd-abcd-abcdefghijkl"
        Type        = "DisconnectParticipant"
        Transitions = {}
        Parameters  = {}
      }
    ]
  })

  tags = [{
    key   = "ModifiedBy"
    value = "AWSCC"
  }]

}
