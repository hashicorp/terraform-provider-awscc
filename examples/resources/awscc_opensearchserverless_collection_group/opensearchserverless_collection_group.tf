resource "awscc_opensearchserverless_collection_group" "example" {
  name        = "example-collection-group"
  description = "Example OpenSearch Serverless collection group for demonstration"

  capacity_limits = {
    max_indexing_capacity_in_ocu = 8
    max_search_capacity_in_ocu   = 8
  }

  standby_replicas = "ENABLED"

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
