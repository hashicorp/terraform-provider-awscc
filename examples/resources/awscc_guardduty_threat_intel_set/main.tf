# Get current AWS account ID
data "aws_caller_identity" "current" {}

# Get current AWS region
data "aws_region" "current" {}

# Get the existing GuardDuty detector
data "aws_guardduty_detector" "current" {}

# Create an S3 bucket for the threat intel list
resource "awscc_s3_bucket" "threat_intel" {
  bucket_name = "threat-intel-set-${data.aws_caller_identity.current.account_id}-${data.aws_region.current.name}"

  public_access_block_configuration = {
    block_public_acls       = true
    block_public_policy     = true
    ignore_public_acls      = true
    restrict_public_buckets = true
  }

  tags = [{
    key   = "Modified By"
    value = "AWSCC"
  }]
}

# Upload a sample threat intel list (using standard AWS provider as AWSCC doesn't have S3 object)
resource "aws_s3_object" "threat_intel_list" {
  bucket  = "threat-intel-set-${data.aws_caller_identity.current.account_id}-${data.aws_region.current.name}"
  key     = "threat-intel.txt"
  content = "8.8.8.8\n8.8.4.4" # Sample IPs for demonstration
}

# Create GuardDuty ThreatIntelSet
resource "awscc_guardduty_threat_intel_set" "example" {
  detector_id = data.aws_guardduty_detector.current.id
  name        = "example-threat-intel-set"
  format      = "TXT"
  location    = "s3://${awscc_s3_bucket.threat_intel.id}/${aws_s3_object.threat_intel_list.key}"
  activate    = true

  tags = [{
    key   = "Modified By"
    value = "AWSCC"
  }]
}

# GuardDuty needs permission to read from the S3 bucket
data "aws_iam_policy_document" "allow_guardduty" {
  statement {
    sid    = "AllowGuardDutyGetObject"
    effect = "Allow"
    principals {
      type        = "Service"
      identifiers = ["guardduty.amazonaws.com"]
    }
    actions = [
      "s3:GetObject"
    ]
    resources = [
      "${awscc_s3_bucket.threat_intel.arn}/*"
    ]
  }
}

# Attach the bucket policy
resource "awscc_s3_bucket_policy" "allow_guardduty" {
  bucket          = awscc_s3_bucket.threat_intel.id
  policy_document = data.aws_iam_policy_document.allow_guardduty.json
}