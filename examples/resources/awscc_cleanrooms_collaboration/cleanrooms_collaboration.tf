resource "awscc_cleanrooms_collaboration" "example" {
  creator_display_name     = "Member1"
  creator_member_abilities = ["CAN_QUERY"]
  description              = "example"
  creator_payment_configuration = {
    query_compute = {
      is_responsible = true
    }
  }
  members = [{
    account_id       = "123456789012"
    member_abilities = ["CAN_RECEIVE_RESULTS"]
    display_name     = "Member2"
    payment_configuration = {
      query_compute = {
        is_responsible = false
      }
    }

    }
  ]
  name             = "example"
  query_log_status = "ENABLED"


  tags = [{
    key   = "Modified By"
    value = "AWSCC"
  }]
}
