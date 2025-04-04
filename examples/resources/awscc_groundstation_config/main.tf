# Get the current AWS account ID
data "aws_caller_identity" "current" {}

# Create an S3 bucket for recording
resource "awscc_s3_bucket" "recording" {
  bucket_name = "aws-groundstation-recording-${data.aws_caller_identity.current.account_id}"
}

# Create IAM role for Ground Station
resource "awscc_iam_role" "ground_station" {
  role_name = "ground-station-recording-role"
  assume_role_policy_document = jsonencode({
    Version = "2012-10-17"
    Statement = [
      {
        Action = "sts:AssumeRole"
        Effect = "Allow"
        Principal = {
          Service = "groundstation.amazonaws.com"
        }
      }
    ]
  })
  managed_policy_arns = [aws_iam_policy.ground_station.arn]
}

# Create IAM policy for Ground Station
resource "aws_iam_policy" "ground_station" {
  name = "ground-station-recording-policy"
  policy = jsonencode({
    Version = "2012-10-17"
    Statement = [
      {
        Effect = "Allow"
        Action = [
          "s3:PutObject",
          "s3:GetObject",
          "s3:ListBucket",
          "s3:GetBucketLocation"
        ]
        Resource = [
          awscc_s3_bucket.recording.arn,
          "${awscc_s3_bucket.recording.arn}/*"
        ]
      }
    ]
  })
}

# Create Ground Station Config
resource "awscc_groundstation_config" "example" {
  name = "example-config"

  config_data = {
    s3_recording_config = {
      bucket_arn = awscc_s3_bucket.recording.arn
      prefix     = "recordings/"
      role_arn   = awscc_iam_role.ground_station.arn
    }
  }

  tags = [{
    key   = "Modified By"
    value = "AWSCC"
  }]
}