# Security Hub v2 configuration with tags
resource "awscc_securityhub_hub_v2" "example" {
  tags = {
    Name        = "example" # Hub identifier
    Environment = "test"    # Environment designation
  }
}
