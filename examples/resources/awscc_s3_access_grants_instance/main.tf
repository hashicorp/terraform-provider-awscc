data "aws_caller_identity" "current" {}
data "aws_region" "current" {}

# IAM Identity Center instance from us-east-1
data "aws_ssoadmin_instances" "example" {
  provider = aws.primary_region
}

resource "awscc_s3_access_grants_instance" "example" {
  identity_center_arn = tolist(data.aws_ssoadmin_instances.example.arns)[0]

  tags = [{
    key   = "Modified By"
    value = "AWSCC"
  }]
}