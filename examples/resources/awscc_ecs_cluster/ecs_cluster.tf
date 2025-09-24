resource "awscc_ecs_cluster" "this" {
  cluster_name = "example_cluster"
  cluster_settings = [{
    name  = "containerInsights"
    value = "enabled"
  }]

  tags = [{
    key   = "Modified By"
    value = "AWSCC"
  }]
}