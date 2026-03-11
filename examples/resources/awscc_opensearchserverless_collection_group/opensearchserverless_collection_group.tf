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

