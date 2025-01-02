data "aws_caller_identity" "current" {}
data "aws_region" "current" {}

# Create an S3 bucket to store the IP set
resource "aws_s3_bucket" "ipset" {
  bucket = "guardduty-ipset-${data.aws_caller_identity.current.account_id}-${data.aws_region.current.name}"
}

# Create the initial IP set file in S3
resource "aws_s3_object" "ipset" {
  bucket  = aws_s3_bucket.ipset.id
  key     = "ipset.txt"
  content = "1.1.1.1/32\n2.2.2.2/32"
}

# Create bucket policy for GuardDuty access
data "aws_iam_policy_document" "bucket_policy" {
  statement {
    actions = ["s3:GetObject"]
    resources = [
      "${aws_s3_bucket.ipset.arn}/*"
    ]
    principals {
      type        = "Service"
      identifiers = ["guardduty.amazonaws.com"]
    }
  }
}

resource "aws_s3_bucket_policy" "ipset" {
  bucket = aws_s3_bucket.ipset.id
  policy = data.aws_iam_policy_document.bucket_policy.json
}

# Get the existing GuardDuty detector
data "aws_guardduty_detector" "main" {}

# Create the IP set
resource "awscc_guardduty_ip_set" "example" {
  detector_id = data.aws_guardduty_detector.main.id
  name        = "example-ipset"
  format      = "TXT"
  location    = "s3://${aws_s3_bucket.ipset.id}/${aws_s3_object.ipset.key}"
  activate    = true

  tags = [{
    key   = "Modified By"
    value = "AWSCC"
  }]
}