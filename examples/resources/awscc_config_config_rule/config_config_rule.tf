resource "awscc_config_config_rule" "example" {
  config_rule_name = "S3_bucket_versioning_enabled"
  description      = "Rule to validate if S3 Bucket versioning is enabled"
  evaluation_modes = [{
    mode = "DETECTIVE"
  }]

  source = {
    owner             = "AWS"
    source_identifier = "S3_BUCKET_VERSIONING_ENABLED"
  }
}