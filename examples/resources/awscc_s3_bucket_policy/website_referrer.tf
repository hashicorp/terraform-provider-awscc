data "aws_caller_identity" "current" {}

resource "awscc_s3_bucket_policy" "example" {
  bucket = awscc_s3_bucket.example.id
  policy_document = jsonencode({
    Version = "2012-10-17"
    Statement = [{
      Action    = ["s3:GetObject"]
      Effect    = "Allow"
      Resource  = "${awscc_s3_bucket.example.arn}/DOC-EXAMPLE-BUCKET/*"
      Principal = "*"
      Condition = {
        StringLike = {
          "aws:Referer" = ["http://www.example.com/*", "http://example.net/*"]
        }
      }
    }]
  })
}

resource "awscc_s3_bucket" "example" {
  public_access_block_configuration = {
    block_public_acls       = true
    block_public_policy     = false
    ignore_public_acls      = true
    restrict_public_buckets = false
  }
}