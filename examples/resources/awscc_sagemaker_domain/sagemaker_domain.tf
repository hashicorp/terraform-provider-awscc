resource "awscc_sagemaker_domain" "example" {
  domain_name = "example"
  auth_mode   = "IAM"
  vpc_id      = awscc_ec2_vpc.example.id
  subnet_ids  = [awscc_ec2_subnet.example.id]

  default_user_settings = {
    execution_role = awscc_iam_role.example.arn
  }
  tags = [
    {
      key   = "Modified By"
      value = "AWSCC"
    }
  ]
}

resource "awscc_iam_role" "example" {
  role_name = "example"
  path      = "/"
  assume_role_policy_document = jsonencode({
    Version = "2012-10-17"
    Statement = [
      {
        Action = "sts:AssumeRole"
        Effect = "Allow"
        Principal = {
          Service = "sagemaker.amazonaws.com"
        }
      },
    ]
  })
  tags = [
    {
      key   = "Modified By"
      value = "AWSCC"
    }
  ]
}


