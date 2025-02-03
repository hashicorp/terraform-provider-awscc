# Get current AWS account ID and region
data "aws_caller_identity" "current" {}
data "aws_region" "current" {}

# Create S3 bucket for CUR reports
resource "awscc_s3_bucket" "cur_bucket" {
  bucket_name = "cur-report-bucket-${data.aws_caller_identity.current.account_id}"

  tags = [{
    key   = "Modified By"
    value = "AWSCC"
  }]
}

# Create bucket policy for CUR to write reports
data "aws_iam_policy_document" "cur_bucket_policy" {
  statement {
    effect = "Allow"
    principals {
      type        = "Service"
      identifiers = ["billingreports.amazonaws.com"]
    }
    actions = [
      "s3:GetBucketAcl",
      "s3:GetBucketPolicy"
    ]
    resources = ["arn:aws:s3:::${awscc_s3_bucket.cur_bucket.bucket_name}"]
  }

  statement {
    effect = "Allow"
    principals {
      type        = "Service"
      identifiers = ["billingreports.amazonaws.com"]
    }
    actions = [
      "s3:PutObject"
    ]
    resources = ["arn:aws:s3:::${awscc_s3_bucket.cur_bucket.bucket_name}/*"]
  }
}

resource "awscc_s3_bucket_policy" "cur_bucket_policy" {
  bucket          = awscc_s3_bucket.cur_bucket.bucket_name
  policy_document = jsonencode(data.aws_iam_policy_document.cur_bucket_policy.json)
}

# Create CUR report definition
resource "awscc_cur_report_definition" "example" {
  report_name                = "cost-usage-report"
  time_unit                  = "HOURLY"
  format                     = "Parquet"
  compression                = "Parquet"
  s3_bucket                  = awscc_s3_bucket.cur_bucket.bucket_name
  s3_prefix                  = "cur-reports"
  s3_region                  = data.aws_region.current.name
  report_versioning          = "OVERWRITE_REPORT"
  refresh_closed_reports     = true
  additional_schema_elements = ["RESOURCES"]
  additional_artifacts       = ["ATHENA"]
}