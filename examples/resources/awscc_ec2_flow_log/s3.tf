data "aws_caller_identity" "current" {}

resource "awscc_ec2_flow_log" "example" {
  log_destination      = awscc_s3_bucket.example.arn
  log_destination_type = "s3"
  traffic_type         = "ALL"
  resource_id          = "vpc-07ddade55bee92f5f"
  resource_type        = "VPC"
  tags = [{
    key   = "Managed By"
    value = "AWSCC"
  }]
}

resource "awscc_s3_bucket" "example" {
  bucket_name = "example-${data.aws_caller_identity.current.account_id}"
  public_access_block_configuration = {
    block_public_acls       = true
    block_public_policy     = true
    ignore_public_acls      = true
    restrict_public_buckets = true
  }
}
