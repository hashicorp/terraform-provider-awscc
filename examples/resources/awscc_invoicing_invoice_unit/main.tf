# Get current AWS account ID
data "aws_caller_identity" "current" {}

# Example Invoice Unit
resource "awscc_invoicing_invoice_unit" "example" {
  name             = "example-invoice-unit"
  description      = "Example Invoice Unit created with AWSCC provider"
  invoice_receiver = data.aws_caller_identity.current.account_id

  rule = {
    linked_accounts = [data.aws_caller_identity.current.account_id]
  }

  tax_inheritance_disabled = false

  resource_tags = [{
    key   = "Modified By"
    value = "AWSCC"
  }]
}