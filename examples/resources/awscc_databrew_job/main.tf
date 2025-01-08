
data "aws_caller_identity" "current" {}
data "aws_region" "current" {}

# Define the IAM role for DataBrew
data "aws_iam_policy_document" "assume_role" {
  statement {
    effect  = "Allow"
    actions = ["sts:AssumeRole"]
    principals {
      type        = "Service"
      identifiers = ["databrew.amazonaws.com"]
    }
  }
}

# Create the IAM role
resource "awscc_iam_role" "databrew_role" {
  role_name                   = "databrew-profile-job-role"
  assume_role_policy_document = data.aws_iam_policy_document.assume_role.json

  tags = [{
    key   = "Modified By"
    value = "AWSCC"
  }]
}

# Define the IAM policy for DataBrew job
data "aws_iam_policy_document" "databrew_policy" {
  statement {
    effect = "Allow"
    actions = [
      "s3:ListBucket",
      "s3:GetBucketLocation"
    ]
    resources = [
      "arn:aws:s3:::${awscc_s3_bucket.example.id}"
    ]
  }

  statement {
    effect = "Allow"
    actions = [
      "s3:GetObject",
      "s3:PutObject",
      "s3:DeleteObject"
    ]
    resources = [
      "arn:aws:s3:::${awscc_s3_bucket.example.id}/*"
    ]
  }

  statement {
    effect = "Allow"
    actions = [
      "logs:CreateLogGroup",
      "logs:CreateLogStream",
      "logs:PutLogEvents"
    ]
    resources = [
      "arn:aws:logs:${data.aws_region.current.name}:${data.aws_caller_identity.current.account_id}:log-group:/aws/databrew/*"
    ]
  }
}

# Attach the policy to the role
resource "awscc_iam_role_policy" "databrew_role_policy" {
  policy_document = data.aws_iam_policy_document.databrew_policy.json
  policy_name     = "databrew-job-policy"
  role_name       = awscc_iam_role.databrew_role.role_name
}

# Create example S3 bucket for outputs
resource "awscc_s3_bucket" "example" {
  bucket_name = "example-databrew-output-${data.aws_caller_identity.current.account_id}"
  tags = [{
    key   = "Modified By"
    value = "AWSCC"
  }]
}

# Upload a sample CSV file
resource "aws_s3_object" "sample" {
  bucket       = awscc_s3_bucket.example.id
  key          = "input/sample.csv"
  content      = "id,name,value\n1,test1,100\n2,test2,200"
  content_type = "text/csv"
}

# Create example dataset
resource "awscc_databrew_dataset" "example" {
  name = "example-dataset"
  input = {
    s3_input_definition = {
      bucket = awscc_s3_bucket.example.id
      key    = "input/sample.csv"
    }
  }
  format = "CSV"
}

# Create the DataBrew Job
resource "awscc_databrew_job" "example" {
  name     = "example-profile-job"
  role_arn = awscc_iam_role.databrew_role.arn
  type     = "PROFILE"

  dataset_name = awscc_databrew_dataset.example.name

  output_location = {
    bucket = awscc_s3_bucket.example.id
    key    = "output/profile-results/"
  }

  tags = [{
    key   = "Modified By"
    value = "AWSCC"
  }]
}