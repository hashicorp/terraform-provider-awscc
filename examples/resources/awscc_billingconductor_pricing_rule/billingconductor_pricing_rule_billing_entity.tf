resource "awscc_billingconductor_pricing_rule" "example_billing_entity" {
  name                = "MarketplaceDiscount"
  
  scope               = "BILLING_ENTITY"
  billing_entity      = "AWS Marketplace"
  type                = "MARKUP"
  modifier_percentage = 5

  tags = [
    {
      key   = "Modified By"
      value = "AWSCC"
    }

  ]
}