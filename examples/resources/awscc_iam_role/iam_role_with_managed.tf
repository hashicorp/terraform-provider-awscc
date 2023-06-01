data "aws_iam_policy_document" "instance_assume_role_policy" {
  statement {
    actions = ["sts:AssumeRole"]

    principals {
      type        = "Service"
      identifiers = ["ec2.amazonaws.com"]
    }
  }
}

resource "aws_iam_policy" "policy_one" {
  name = "policy_one"

  policy = jsonencode({
    Version = "2012-10-17"
    Statement = [
      {
        Action   = ["ec2:Describe*"]
        Effect   = "Allow"
        Resource = "*"
      },
    ]
  })
}

resource "aws_iam_policy" "policy_two" {
  name = "policy_two"

  policy = jsonencode({
    Version = "2012-10-17"
    Statement = [
      {
        Action   = ["s3:ListAllMyBuckets", "s3:ListBucket", "s3:HeadBucket"]
        Effect   = "Allow"
        Resource = "*"
      },
    ]
  })
}


resource "awscc_iam_role" "main" {
  role_name                   = "sample_iam_role"
  description                 = "This is a sample IAM role"
  assume_role_policy_document = data.aws_iam_policy_document.instance_assume_role_policy.json
  managed_policy_arns         = [aws_iam_policy.policy_one.arn, aws_iam_policy.policy_two.arn]
  path                        = "/"
  tags = [
    {
      key   = "Name"
      value = "Sample IAM Role"
    },
    {
      key   = "Environment"
      value = "Development"
    },
    {
      key   = "Modified By"
      value = "AWSCC"
    }
  ]
}