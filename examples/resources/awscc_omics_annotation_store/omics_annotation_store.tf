# Create a KMS key for encryption
resource "aws_kms_key" "omics_key" {
  description             = "KMS key for Omics resources"
  deletion_window_in_days = 7
}

# Create reference store using AWSCC
resource "awscc_omics_reference_store" "example" {
  name        = "example-reference-store"
  description = "Example reference store"
}

resource "awscc_omics_annotation_store" "example" {
  name         = "example-annotation-store"
  description  = "Example annotation store created with Terraform"
  store_format = "VCF"

  reference = {
    reference_arn = awscc_omics_reference_store.example.arn
  }

  sse_config = {
    type    = "KMS"
    key_arn = aws_kms_key.omics_key.arn
  }

  tags = {
    Environment = "example"
    Name        = "example-annotation-store"
  }
}
