# Data sources for region and caller identity
data "aws_region" "current" {}
data "aws_caller_identity" "current" {}

# Create an AWS SSM Association for running AWS-RunPowerShellScript
resource "awscc_ssm_association" "example" {
  name             = "AWS-RunPowerShellScript"
  association_name = "example-powershell-association"

  targets = [{
    key    = "tag:Environment"
    values = ["Production"]
  }]

  parameters = {
    commands = ["Get-Service"]
  }

  schedule_expression = "cron(0 0 * * ? *)"

  output_location = {
    s3_location = {
      output_s3_bucket_name = awscc_s3_bucket.output.id
      output_s3_key_prefix  = "ssm-output/"
      output_s3_region      = data.aws_region.current.name
    }
  }

  max_concurrency     = "50%"
  max_errors          = "25%"
  compliance_severity = "MEDIUM"
}

# Create S3 bucket for output
resource "awscc_s3_bucket" "output" {
  bucket_name = "ssm-association-output-${data.aws_caller_identity.current.account_id}-${data.aws_region.current.name}"

  public_access_block_configuration = {
    block_public_acls       = true
    block_public_policy     = true
    ignore_public_acls      = true
    restrict_public_buckets = true
  }

  ownership_controls = {
    rules = [{
      object_ownership = "BucketOwnerPreferred"
    }]
  }

  tags = [{
    key   = "Modified By"
    value = "AWSCC"
  }]
}

# Create bucket policy for SSM
data "aws_iam_policy_document" "bucket_policy" {
  statement {
    effect = "Allow"
    principals {
      type        = "Service"
      identifiers = ["ssm.amazonaws.com"]
    }
    actions = [
      "s3:PutObject"
    ]
    resources = [
      "${awscc_s3_bucket.output.arn}/*"
    ]
    condition {
      test     = "StringEquals"
      variable = "aws:SourceAccount"
      values   = [data.aws_caller_identity.current.account_id]
    }
    condition {
      test     = "StringEquals"
      variable = "aws:SourceRegion"
      values   = [data.aws_region.current.name]
    }
  }
}

resource "awscc_s3_bucket_policy" "allow_ssm" {
  bucket          = awscc_s3_bucket.output.id
  policy_document = data.aws_iam_policy_document.bucket_policy.json
}