# Create a RAM resource share
resource "awscc_ram_resource_share" "example" {
  name                      = "test-ram-share"
  allow_external_principals = true

  # Add VPC sharing permission
  permission_arns = [
    "arn:aws:ram::aws:permission/AWSRAMDefaultPermissionSubnet"
  ]

  # Example principal - replace with actual AWS account ID in your environment
  principals = [
    "123456789012"
  ]

  # Tags as per AWSCC format
  tags = [{
    key   = "Modified By"
    value = "AWSCC"
  }]
}