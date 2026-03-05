# OpenSearch Serverless Collection Group
resource "awscc_opensearchserverless_collection_group" "example" {
  name             = "example-collection-group"
  standby_replicas = "ENABLED"
  description      = "Example OpenSearch Serverless collection group"

  capacity_limits = {
    max_indexing_capacity_in_ocu = 16
    max_search_capacity_in_ocu   = 16
    min_indexing_capacity_in_ocu = 2
    min_search_capacity_in_ocu   = 2
  }

  tags = [
    {
      key   = "Name"
      value = "example-collection-group"
    },
    {
      key   = "Environment"
      value = "example"
    }
  ]
}

output "collection_group_arn" {
  description = "ARN of the OpenSearch Serverless collection group"
  value       = awscc_opensearchserverless_collection_group.example.arn
}

output "collection_group_id" {
  description = "ID of the OpenSearch Serverless collection group"
  value       = awscc_opensearchserverless_collection_group.example.collection_group_id
}