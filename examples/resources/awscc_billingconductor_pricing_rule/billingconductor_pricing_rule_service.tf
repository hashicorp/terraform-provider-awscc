resource "awscc_billingconductor_pricing_rule" "example_service" {
  name = "S3Discount"

  scope               = "SERVICE"
  service             = "AmazonS3"
  type                = "DISCOUNT"
  modifier_percentage = 5

  tags = [
    {
      key   = "Modified By"
      value = "AWSCC"
    }

  ]
}