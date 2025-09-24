resource "awscc_bedrock_data_source" "example" {
  name              = "example"
  knowledge_base_id = awscc_bedrock_knowledge_base.example.knowledge_base_id
  description       = "Example datasource"

  data_source_configuration = {
    s3_configuration = {
      bucket_arn         = var.bucket_arn
      inclusion_prefixes = ["example"]
    }
    type = "S3"
  }

  server_side_encryption_configuration = {
    kms_key_arn = var.kms_key_arn
  }
}

variable "bucket_arn" {
  type = string
}

variable "kms_key_arn" {
  type = string
}