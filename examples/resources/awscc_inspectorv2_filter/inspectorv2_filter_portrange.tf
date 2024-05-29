resource "awscc_inspectorv2_filter" "example" {
  name          = "example"
  description   = "Suppression rule filtered on port range in a specific account"
  filter_action = "NONE"
  filter_criteria = {
    aws_account_id = [{
      comparison = "EQUALS"
      value      = data.aws_caller_identity.current.account_id
    }]
    port_range = [{
      begin_inclusive = 0
      end_inclusive   = 65535
    }]
  }

}