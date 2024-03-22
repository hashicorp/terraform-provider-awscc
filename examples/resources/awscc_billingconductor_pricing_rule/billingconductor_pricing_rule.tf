resource "awscc_billingconductor_pricing_rule" "example" {
  name                = "TestPricingRule-{random_id.test}"
  description         = "Mark up everything by 10%."
  scope               = "GLOBAL"
  type                = "MARKUP"
  modifier_percentage = 10

  tags = [
    {
      key   = "Modified By"
      value = "AWSCC"
    }

  ]
}