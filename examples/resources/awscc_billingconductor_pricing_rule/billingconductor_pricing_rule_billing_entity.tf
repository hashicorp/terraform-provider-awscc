resource "awscc_billingconductor_pricing_rule" "example" {
  name                = "test-markup-by-billing_entitiy_marketplace"
  description         = "Markup 10% if use market place"
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