resource "awscc_iam_instance_profile" "example" {
  roles                 = [awscc_iam_role.example.role_name]
  instance_profile_name = "dev_profile"
  path                  = "/compute/"
}

resource "awscc_iam_role_policy" "example" {
  policy_name = "minimal_ssm_policy"
  role_name   = awscc_iam_role.example.role_name
  policy_document = jsonencode({
    Version = "2012-10-17"
    Statement = [
      {
        Action = [
          "ssm:UpdateInstanceInformation",
          "ssmmessages:CreateControlChannel",
          "ssmmessages:CreateDataChannel",
          "ssmmessages:OpenControlChannel",
          "ssmmessages:OpenDataChannel"
        ]
        Effect   = "Allow"
        Resource = "*"
      }
    ]
  })
}

resource "awscc_iam_role" "example" {
  role_name = "dev_compute_role"
  assume_role_policy_document = jsonencode({
    Version = "2012-10-17"
    Statement = [
      {
        Action = "sts:AssumeRole"
        Effect = "Allow"
        Principal = {
          Service = "ec2.amazonaws.com"
        }
      },
    ]
  })
}
