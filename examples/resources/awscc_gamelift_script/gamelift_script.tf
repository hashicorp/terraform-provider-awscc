resource "awscc_gamelift_script" "example" {
  name    = "example-script"
  version = "v1.0.0"

  storage_location = {
    bucket   = "your-s3-bucket"
    key      = "your-s3-key"
    role_arn = awscc_iam_role.example.arn
  }
}

resource "awscc_iam_role" "example" {
  role_name                   = "gamelift-s3-access"
  description                 = "This IAM role grants Amazon GameLift access to the S3 bucket containing script files."
  assume_role_policy_document = data.aws_iam_policy_document.instance_assume_role_policy.json
  managed_policy_arns         = [aws_iam_policy.example.arn]
  max_session_duration        = 7200
  path                        = "/"
  tags = [
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

data "aws_iam_policy_document" "instance_assume_role_policy" {
  statement {
    actions = ["sts:AssumeRole"]

    principals {
      type        = "Service"
      identifiers = ["gamelift.amazonaws.com"]
    }
  }
}

resource "aws_iam_policy" "example" {
  name = "gamelift-s3-access-policy"

  policy = jsonencode({
    Version = "2012-10-17"
    Statement = [
      {
        Effect   = "Allow"
        Action   = ["s3:*"]
        Resource = "*"
      },
    ]
  })
}
