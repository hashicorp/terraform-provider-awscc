resource "awscc_inspectorv2_filter" "example" {
  name          = "example"
  description   = "Suppression rule using account filter"
  filter_action = "SUPPRESS"
  filter_criteria = {
    aws_account_id = [{
      comparison = "EQUALS"
      value      = data.aws_caller_identity.current.account_id
    }]
  }
}