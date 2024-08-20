resource "awscc_billingconductor_billing_group" "example" {
  name        = "example"
  description = "Example description"
  computation_preference = {
    pricing_plan_arn = awscc_billingconductor_pricing_plan.example.arn
  }
  primary_account_id = "123456789012"
  account_grouping = {
    linked_account_ids = ["123456789012"]
  }
  tags = [
    {
      key   = "Modified By"
      value = "AWSCC"
    }
  ]
}