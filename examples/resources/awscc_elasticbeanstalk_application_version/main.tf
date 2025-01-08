data "aws_caller_identity" "current" {}
data "aws_region" "current" {}

# Create an S3 bucket for the application source
resource "awscc_s3_bucket" "app_bucket" {
  bucket_name = "beanstalk-app-source-${data.aws_caller_identity.current.account_id}-${data.aws_region.current.name}"
  tags = [{
    key   = "Modified By"
    value = "AWSCC"
  }]
}

# Policy to allow ElasticBeanstalk to access the S3 bucket
data "aws_iam_policy_document" "allow_beanstalk" {
  statement {
    effect = "Allow"
    principals {
      type        = "Service"
      identifiers = ["elasticbeanstalk.amazonaws.com"]
    }
    actions = [
      "s3:GetObject",
      "s3:GetObjectVersion"
    ]
    resources = [
      "arn:aws:s3:::${awscc_s3_bucket.app_bucket.id}/*"
    ]
  }
}

# Attach the bucket policy
resource "awscc_s3_bucket_policy" "app_bucket_policy" {
  bucket          = awscc_s3_bucket.app_bucket.id
  policy_document = jsonencode(data.aws_iam_policy_document.allow_beanstalk.json)
}

# Create a sample application file
resource "aws_s3_object" "app_source" {
  bucket = awscc_s3_bucket.app_bucket.id
  key    = "sample-app.zip"
  source = "sample-app.zip"
}

# Create the Elastic Beanstalk application
resource "awscc_elasticbeanstalk_application" "example" {
  application_name = "example-app"
  description      = "Example Elastic Beanstalk Application"
}

# Create the application version
resource "awscc_elasticbeanstalk_application_version" "example" {
  application_name = awscc_elasticbeanstalk_application.example.application_name
  description      = "Example Application Version"
  source_bundle = {
    s3_bucket = awscc_s3_bucket.app_bucket.id
    s3_key    = aws_s3_object.app_source.key
  }
}