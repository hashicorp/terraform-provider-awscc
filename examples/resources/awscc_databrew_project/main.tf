data "aws_caller_identity" "current" {}
data "aws_region" "current" {}

# Create an S3 bucket for the DataBrew dataset
resource "awscc_s3_bucket" "databrew_input" {
  bucket_name = "databrew-input-${data.aws_caller_identity.current.account_id}-${data.aws_region.current.name}"

  tags = [{
    key   = "Modified By"
    value = "AWSCC"
  }]
}

# Upload a sample CSV file to the bucket
resource "aws_s3_object" "sample_data" {
  bucket  = awscc_s3_bucket.databrew_input.id
  key     = "input/sample.csv"
  content = "id,name,age\n1,John,30\n2,Jane,25"
}

# Create the DataBrew dataset
resource "awscc_databrew_dataset" "example" {
  name = "example-dataset"

  input = {
    s3_input_definition = {
      bucket = awscc_s3_bucket.databrew_input.id
      key    = aws_s3_object.sample_data.key
    }
  }

  format = "CSV"
  format_options = {
    csv = {
      delimiter  = ","
      header_row = true
    }
  }

  tags = [{
    key   = "Modified By"
    value = "AWSCC"
  }]
}

# Create an IAM role to allow Databrew access to S3 Bucket
resource "awscc_iam_role" "example" {
  assume_role_policy_document = jsonencode({
    Version = "2012-10-17"
    Statement = [
      {
        Action = "sts:AssumeRole"
        Effect = "Allow"
        Principal = {
          Service = "databrew.amazonaws.com"
        }
      },
    ]
  })
}

resource "awscc_iam_role_policy" "example" {
  role_name   = awscc_iam_role.example.role_name
  policy_name = "example-policy"
  policy_document = jsonencode({
    Version = "2012-10-17"
    Statement = [
      {
        Action = [
          "s3:ListBucket",
          "s3:GetObject"
        ]
        Effect = "Allow"
        Resource = [
          awscc_s3_bucket.databrew_input.arn,
          "${awscc_s3_bucket.databrew_input.arn}/*"
        ]
      }
    ]
  })
}

# Create a DataBrew project
resource "awscc_databrew_project" "example" {
  dataset_name = awscc_databrew_dataset.example.name
  name         = "example-project"
  recipe_name  = "example-recipe"
  role_arn     = awscc_iam_role.example.arn
}
