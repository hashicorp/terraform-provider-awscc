// Create new Amazon GameLift Build
resource "awscc_gamelift_build" "example" {
  name             = "example-build"
  version          = "1"              // (Optional) - used for versioning your GameLift Builds
  operating_system = "AMAZON_LINUX_2" // (Required) - valid values are "WINDOWS_2012", "AMAZON_LINUX", or "AMAZON_LINUX_2"

  // Required
  storage_location = {
    bucket   = "your-s3-bucket"           // Name of the S3 bucket your build files are stored in
    key      = "your-s3-key"              // Name of the .zip file containing your build files
    role_arn = awscc_iam_role.example.arn // ARN of the AWS IAM Role that allows Amazon GameLift to access your S3 Bucket
  }
}

// Create IAM Role that allows GameLift to access your S3 bucket containing build files
resource "awscc_iam_role" "example" {
  role_name                   = "gamelift-s3-access"
  description                 = "This IAM role grants Amazon GameLift access to the S3 bucket containing build files"
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

// Creat Trust Relationshop to allow GameLift to assume the IAM Role
data "aws_iam_policy_document" "instance_assume_role_policy" {
  statement {
    actions = ["sts:AssumeRole"]

    principals {
      type        = "Service"
      identifiers = ["gamelift.amazonaws.com"]
    }
  }
}

// Create IAM Customer Managed Policy for S3 access
resource "aws_iam_policy" "example" {
  name = "gamelift-s3-access-policy"

  policy = jsonencode({
    Version = "2012-10-17"
    Statement = [
      {
        Effect   = "Allow"
        Action   = ["s3:*"] // IMPORTANT: Example meant for testing purposes only. Restrict these permissions further for enhanced security.
        Resource = "*"      // IMPORTANT: Example meant for testing purposes only. Restrict these permissions further for enhanced security.
      },
    ]
  })
}
