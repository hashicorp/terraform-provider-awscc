# Get current AWS account ID
data "aws_caller_identity" "current" {}

# First, we need to create a cleanrooms collaboration
resource "awscc_cleanrooms_collaboration" "example" {
  name                 = "example-collaboration"
  description          = "Example collaboration for membership"
  creator_display_name = "Creator"
  creator_member_abilities = [
    "CAN_QUERY",
    "CAN_RECEIVE_RESULTS"
  ]
  members = [{
    account_id       = "123456789012" # Replace with actual member account ID
    display_name     = "Member 1"
    member_abilities = ["CAN_QUERY", "CAN_RECEIVE_RESULTS"]
  }]
  query_log_status = "ENABLED"
  tags = [{
    key   = "Modified By"
    value = "AWSCC"
  }]
}

# Create an S3 bucket for query results
resource "awscc_s3_bucket" "query_results" {
  bucket_name = "cleanrooms-query-results-${data.aws_caller_identity.current.account_id}"
  tags = [{
    key   = "Modified By"
    value = "AWSCC"
  }]
}

# Create IAM role for Clean Rooms
resource "awscc_iam_role" "cleanrooms_role" {
  role_name = "cleanrooms-query-role"
  assume_role_policy_document = jsonencode({
    Version = "2012-10-17"
    Statement = [
      {
        Action = "sts:AssumeRole"
        Effect = "Allow"
        Principal = {
          Service = "cleanrooms.amazonaws.com"
        }
      }
    ]
  })
  tags = [{
    key   = "Modified By"
    value = "AWSCC"
  }]
}

# IAM policy for Clean Rooms to access S3
resource "awscc_iam_role_policy" "cleanrooms_s3_policy" {
  policy_name = "cleanrooms-s3-policy"
  role_name   = awscc_iam_role.cleanrooms_role.role_name
  policy_document = jsonencode({
    Version = "2012-10-17"
    Statement = [
      {
        Effect = "Allow"
        Action = [
          "s3:PutObject",
          "s3:GetObject",
          "s3:ListBucket"
        ]
        Resource = [
          "arn:aws:s3:::${awscc_s3_bucket.query_results.bucket_name}",
          "arn:aws:s3:::${awscc_s3_bucket.query_results.bucket_name}/*"
        ]
      }
    ]
  })
}

# Create the Clean Rooms membership
resource "awscc_cleanrooms_membership" "example" {
  collaboration_identifier = awscc_cleanrooms_collaboration.example.id
  query_log_status         = "ENABLED"

  default_result_configuration = {
    role_arn = awscc_iam_role.cleanrooms_role.arn
    output_configuration = {
      s3 = {
        bucket             = awscc_s3_bucket.query_results.bucket_name
        key_prefix         = "results/"
        result_format      = "CSV"
        single_file_output = true
      }
    }
  }

  tags = [{
    key   = "Modified By"
    value = "AWSCC"
  }]
}