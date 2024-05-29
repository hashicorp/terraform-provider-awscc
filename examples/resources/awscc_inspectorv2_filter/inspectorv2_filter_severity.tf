resource "awscc_inspectorv2_filter" "example" {
  name          = "example"
  description   = "Suppression rule filtered on severity  matching Critical or High in a specific account"
  filter_action = "NONE"
  filter_criteria = {
    aws_account_id = [{
      comparison = "EQUALS"
      value      = data.aws_caller_identity.current.account_id
    }]
    severity = [{
      comparison = "EQUALS"
      value      = "Critical"
      },
      {
        comparison = "EQUALS"
        value      = "High"
      }
    ]
  }

}