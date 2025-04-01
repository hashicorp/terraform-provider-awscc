# Get current AWS account ID
data "aws_caller_identity" "current" {}

# Create S3 bucket for access grants location
resource "awscc_s3_bucket" "example" {
  bucket_name = "example-access-grants-bucket-${data.aws_caller_identity.current.account_id}"

  tags = [{
    key   = "Modified By"
    value = "AWSCC"
  }]
}

# IAM Role for the Access Grants Instance
resource "awscc_iam_role" "access_grants_instance_role" {
  role_name = "s3-access-grants-instance-role"
  assume_role_policy_document = jsonencode({
    Version = "2012-10-17"
    Statement = [
      {
        Action = "sts:AssumeRole"
        Effect = "Allow"
        Principal = {
          Service = "s3.amazonaws.com"
        }
      }
    ]
  })

  tags = [{
    key   = "Modified By"
    value = "AWSCC"
  }]
}

# Create Access Grants Instance Role Policy
resource "aws_iam_policy" "access_grants_instance_policy" {
  name = "s3-access-grants-instance-policy"
  policy = jsonencode({
    Version = "2012-10-17"
    Statement = [
      {
        Effect = "Allow"
        Action = [
          "s3:*"
        ]
        Resource = "*"
      }
    ]
  })
}

# Attach policy to instance role via managed_policy_arns
resource "awscc_iam_role" "access_grants_instance_role_with_policy" {
  role_name = "s3-access-grants-instance-role-with-policy"
  assume_role_policy_document = jsonencode({
    Version = "2012-10-17"
    Statement = [
      {
        Action = "sts:AssumeRole"
        Effect = "Allow"
        Principal = {
          Service = "s3.amazonaws.com"
        }
      }
    ]
  })

  managed_policy_arns = [aws_iam_policy.access_grants_instance_policy.arn]

  tags = [{
    key   = "Modified By"
    value = "AWSCC"
  }]
}

# Create Access Grants Instance
resource "awscc_s3_access_grants_instance" "example" {
  tags = [{
    key   = "Modified By"
    value = "AWSCC"
  }]
}

# IAM Role for access grants location
resource "awscc_iam_role" "access_grants_role" {
  role_name = "s3-access-grants-location-role"
  assume_role_policy_document = jsonencode({
    Version = "2012-10-17"
    Statement = [
      {
        Action = "sts:AssumeRole"
        Effect = "Allow"
        Principal = {
          Service = "s3.amazonaws.com"
        }
      }
    ]
  })

  path                = "/"
  managed_policy_arns = [aws_iam_policy.access_grants_policy.arn]

  tags = [{
    key   = "Modified By"
    value = "AWSCC"
  }]
}

# Create IAM policy for access grants
resource "aws_iam_policy" "access_grants_policy" {
  name = "s3-access-grants-location-policy"
  policy = jsonencode({
    Version = "2012-10-17"
    Statement = [
      {
        Effect = "Allow"
        Action = [
          "s3:GetObject",
          "s3:PutObject",
          "s3:DeleteObject",
          "s3:ListBucket"
        ]
        Resource = [
          "arn:aws:s3:::${awscc_s3_bucket.example.bucket_name}",
          "arn:aws:s3:::${awscc_s3_bucket.example.bucket_name}/*"
        ]
      }
    ]
  })
}

# S3 Access Grants Location
resource "awscc_s3_access_grants_location" "example" {
  depends_on = [awscc_s3_access_grants_instance.example]

  location_scope = "s3://${awscc_s3_bucket.example.bucket_name}/reports/"
  iam_role_arn   = awscc_iam_role.access_grants_role.arn

  tags = [{
    key   = "Modified By"
    value = "AWSCC"
  }]
}