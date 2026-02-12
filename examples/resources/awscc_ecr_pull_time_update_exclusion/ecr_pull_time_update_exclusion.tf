resource "awscc_iam_role" "example_role" {
  role_name = "example-ecr-pull-time-exclusion-role"

  assume_role_policy_document = jsonencode({
    Version = "2012-10-17"
    Statement = [
      {
        Action = "sts:AssumeRole"
        Effect = "Allow"
        Principal = {
          Service = "ec2.amazonaws.com"
        }
      }
    ]
  })

  tags = [
    {
      key   = "Name"
      value = "example-ecr-pull-time-exclusion-role"
    },
    {
      key   = "Environment"
      value = "example"
    }
  ]
}

resource "awscc_ecr_pull_time_update_exclusion" "example" {
  principal_arn = awscc_iam_role.example_role.arn
}
