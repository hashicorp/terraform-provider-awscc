# Create an IAM user first for testing
resource "awscc_iam_user" "example" {
  user_name = "example-mfa-user"
  path      = "/"
  tags = [{
    key   = "Modified By"
    value = "AWSCC"
  }]
}

# Create a virtual MFA device
resource "awscc_iam_virtual_mfa_device" "example" {
  virtual_mfa_device_name = "example-mfa-device"
  users                   = [awscc_iam_user.example.user_name]
  path                    = "/"
  tags = [{
    key   = "Modified By"
    value = "AWSCC"
  }]
}