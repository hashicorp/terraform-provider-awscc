resource "awscc_ecs_cluster_capacity_provider_associations" "example" {
  cluster            = awscc_ecs_cluster.example.name
  capacity_providers = [awscc_ecs_capacity_provider.example.id]
}

resource "awscc_ecs_capacity_provider" "example" {
  name = "example"

  auto_scaling_group_provider = {
    auto_scaling_group_arn         = aws_autoscaling_group.example.arn
    managed_termination_protection = "ENABLED"

    managed_scaling = {
      maximum_scaling_step_size = 1000
      minimum_scaling_step_size = 1
      status                    = "ENABLED"
      target_capacity           = 10
    }
  }
  lifecycle {
    ignore_changes = [auto_scaling_group_provider.auto_scaling_group_arn]
  }
}
