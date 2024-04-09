resource "awscc_billingconductor_pricing_rule" "example" {
  name                = "Markup10percent"
  
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