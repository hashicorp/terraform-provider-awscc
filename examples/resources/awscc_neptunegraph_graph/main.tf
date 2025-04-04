# Create Neptune Graph
resource "awscc_neptunegraph_graph" "example" {
  graph_name           = "example-graph-test-20250102"
  provisioned_memory   = 16
  deletion_protection  = false
  public_connectivity = false
  replica_count       = 1

  vector_search_configuration = {
    vector_search_dimension = 128
  }

  tags = [
    {
      key   = "Environment"
      value = "Development"
    },
    {
      key   = "ModifiedBy"
      value = "AWSCC"
    }
  ]
}