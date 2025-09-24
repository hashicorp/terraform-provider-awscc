resource "awscc_wisdom_assistant" "example" {
  name        = "example"
  description = "Example assistant"
  type        = "AGENT"
  server_side_encryption_configuration = {
    kms_key_id = awscc_kms_key.example.arn
  }

  tags = [{
    key   = "ModifiedBy"
    value = "AWSCC"
  }]

}