resource "awscc_neptunegraph_graph" "example" {
  graph_name          = "example-graph"
  provisioned_memory  = 16
  deletion_protection = false
  public_connectivity = false
  replica_count       = 1
  vector_search_configuration = {
    vector_search_dimension = 1536
  }

  tags = {
    Modified_By = "AWSCC"
  }
}

resource "awscc_neptunegraph_graph_snapshot" "example" {
  graph_identifier = awscc_neptunegraph_graph.example.graph_id
  snapshot_name    = "example-snapshot"

  tags = {
    Modified_By = "AWSCC"
  }
}

# Export the graph snapshot ID
output "graph_snapshot_id" {
  description = "ID of the Neptune Analytics graph snapshot"
  value       = awscc_neptunegraph_graph_snapshot.example.snapshot_id
}

# Export the snapshot ARN
output "snapshot_arn" {
  description = "ARN of the Neptune Analytics graph snapshot"
  value       = awscc_neptunegraph_graph_snapshot.example.snapshot_arn
}
