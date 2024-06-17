resource "awscc_securityhub_product_subscription" "example" {
  # Alert Logic – SIEMless Threat Management
  product_arn = "arn:aws:securityhub:us-east-1:733251395267:product/alertlogic/althreatmanagement"
}
