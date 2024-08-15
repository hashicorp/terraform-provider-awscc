resource "awscc_billingconductor_custom_line_item" "example" {
  billing_group_arn = awscc_billingconductor_billing_group.example.arn
  name              = "example"
  description       = "Example description"
  account_id        = "123456789012"
  custom_line_item_charge_details = {
    flat = {
      charge_value = 1
    }
    type = "FEE"
  }
  billing_period_range = {
    inclusive_start_billing_period = "2024-07"
  }
  tags = [
    {
      key   = "Modified By"
      value = "AWSCC"
    }
  ]

}
