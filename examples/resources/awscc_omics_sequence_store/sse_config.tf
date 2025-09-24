resource "awscc_omics_sequence_store" "example" {
  name        = "example"
  description = "example"

  sse_config = {
    type    = "KMS"
    key_arn = awscc_kms_key.example.arn
  }
}