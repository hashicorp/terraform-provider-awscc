# IAM Identity Center instance
data "aws_ssoadmin_instances" "example" {
}

resource "awscc_s3_access_grants_instance" "example" {
  identity_center_arn = tolist(data.aws_ssoadmin_instances.example.arns)[0]

  tags = [{
    key   = "Modified By"
    value = "AWSCC"
  }]
}