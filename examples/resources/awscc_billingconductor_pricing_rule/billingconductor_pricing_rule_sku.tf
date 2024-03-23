resource "awscc_billingconductor_pricing_rule" "example_sku" {
  name        = "DiscountEC2_T2Micro_LinuxUnix"
  description = "5% Discount for t2.micro on Linux/Unix in Singapore region"

  scope      = "SKU"
  service    = "AmazonEC2"
  usage_type = "APS1-BoxUsage:t2.medium"
  operation  = "RunInstances"

  type                = "DISCOUNT"
  modifier_percentage = 5

  tags = [
    {
      key   = "Modified By"
      value = "AWSCC"
    }

  ]
}