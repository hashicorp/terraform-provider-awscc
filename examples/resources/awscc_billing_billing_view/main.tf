resource "awscc_billing_billing_view" "example" {
  name         = "example-billing-view"
  description  = "Example billing view for cost analysis"
  source_views = ["arn:aws:billing::123456789012:billingview/source-view"] # Set this to the right arn and linked account

  data_filter_expression = {
    dimensions = {
      key    = "LINKED_ACCOUNT"
      values = ["123456789012"]
    }
  }

  tags = [{
    key   = "Modified By"
    value = "AWSCC"
  }]
}
