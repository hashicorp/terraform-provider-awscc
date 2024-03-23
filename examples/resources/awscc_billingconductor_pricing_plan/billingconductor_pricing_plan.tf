
resource "awscc_billingconductor_pricing_plan" "example" {
  name              = "HR-Rates"
  pricing_rule_arns = [awscc_billingconductor_pricing_rule.example.arn]

  tags = [
    {
      key   = "Modified By"
      value = "AWSCC"
    }
  ]
}