resource "awscc_bedrock_knowledge_base" "example" {
  name        = "example"
  description = "Example Knowledge base"
  role_arn    = var.kb_role_arn

  storage_configuration = {
    type = "OPENSEARCH_SERVERLESS"
    opensearch_serverless_configuration = {
      collection_arn    = var.collection_arn
      vector_index_name = var.vector_index_name
      field_mapping = {
        metadata_field = var.metadata_field
        text_field     = var.text_field
        vector_field   = var.vector_field
      }
    }
  }
  knowledge_base_configuration = {
    type = "VECTOR"
    vector_knowledge_base_configuration = {
      embedding_model_arn = "arn:aws:bedrock:us-east-1::foundation-model/amazon.titan-embed-text-v1"
    }
  }
}

variable "vector_index_name" {
  type = string
}

variable "metadata_field" {
  type = string
}

variable "text_field" {
  type = string
}

variable "vector_field" {
  type = string
}

variable "collection_arn" {
  type = string
}

variable "kb_role_arn" {
  type = string
}
