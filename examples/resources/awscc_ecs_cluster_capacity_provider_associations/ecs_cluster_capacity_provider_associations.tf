resource "awscc_ecs_cluster_capacity_provider_associations" "example" {
  cluster            = awscc_ecs_cluster.example.name
  capacity_providers = ["FARGATE", "FARGATE_SPOT"]

  default_capacity_provider_strategy = [{
    capacity_provider = "FARGATE"
    base              = 1
    weight            = 100

  }]
}