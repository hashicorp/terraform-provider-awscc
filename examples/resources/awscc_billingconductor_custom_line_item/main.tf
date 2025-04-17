data "aws_caller_identity" "current" {}

# Create the custom line item
resource "awscc_billingconductor_custom_line_item" "example" {
  name        = "example-custom-line-item"
  description = "Example custom line item for demonstration"

  # The ARN must be an existing billing group
  billing_group_arn = "arn:aws:billingconductor::${data.aws_caller_identity.current.account_id}:billinggroup/${data.aws_caller_identity.current.account_id}"

  # Optional: specify the account to charge
  account_id = data.aws_caller_identity.current.account_id

  # Optional: specify billing period range
  billing_period_range = {
    inclusive_start_billing_period = "2024-12"
    exclusive_end_billing_period   = "2025-01"
  }

  # Specify charge details
  custom_line_item_charge_details = {
    type = "FEE"
    flat = {
      charge_value = 100.0
    }
  }

  tags = [{
    key   = "Modified By"
    value = "AWSCC"
  }]
}