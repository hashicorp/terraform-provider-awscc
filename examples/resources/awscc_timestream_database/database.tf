resource "aws_kms_key" "this" {
  description = "Timestream KMS Key"
}

resource "awscc_timestream_database" "this" {
  database_name = "MyTimestreamDB"
  kms_key_id    = aws_kms_key.this.key_id
  tags = [{
    key   = "Managed By"
    value = "AWSCC"
  }]
}
