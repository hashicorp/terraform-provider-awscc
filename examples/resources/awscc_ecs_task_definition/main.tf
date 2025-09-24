# Get current AWS region
data "aws_region" "current" {}

# Task execution role and policy
resource "awscc_iam_role" "execution_role" {
  assume_role_policy_document = jsonencode({
    Version = "2012-10-17"
    Statement = [
      {
        Action = "sts:AssumeRole"
        Effect = "Allow"
        Principal = {
          Service = "ecs-tasks.amazonaws.com"
        }
      }
    ]
  })
  managed_policy_arns = [
    "arn:aws:iam::aws:policy/service-role/AmazonECSTaskExecutionRolePolicy"
  ]
  description = "ECS task execution role"
  path        = "/"
  role_name   = "ecs-task-execution-role-example"
  tags = [{
    key   = "Modified By"
    value = "AWSCC"
  }]
}

# Task role and policy
resource "awscc_iam_role" "task_role" {
  assume_role_policy_document = jsonencode({
    Version = "2012-10-17"
    Statement = [
      {
        Action = "sts:AssumeRole"
        Effect = "Allow"
        Principal = {
          Service = "ecs-tasks.amazonaws.com"
        }
      }
    ]
  })
  description = "ECS task role"
  path        = "/"
  role_name   = "ecs-task-role-example"
  tags = [{
    key   = "Modified By"
    value = "AWSCC"
  }]
}

# Task definition
resource "awscc_ecs_task_definition" "example" {
  family = "example-task-definition"
  requires_compatibilities = [
    "FARGATE"
  ]
  execution_role_arn = awscc_iam_role.execution_role.arn
  task_role_arn      = awscc_iam_role.task_role.arn
  cpu                = "256"
  memory             = "512"
  network_mode       = "awsvpc"

  runtime_platform = {
    operating_system_family = "LINUX"
    cpu_architecture        = "X86_64"
  }

  container_definitions = [
    {
      name      = "nginx"
      image     = "nginx:latest"
      essential = true

      port_mappings = [
        {
          container_port = 80
          protocol       = "tcp"
        }
      ]

      environment = [
        {
          name  = "ENVIRONMENT"
          value = "example"
        }
      ]

      log_configuration = {
        log_driver = "awslogs"
        options = {
          "awslogs-group"         = "/ecs/example-task-definition"
          "awslogs-region"        = data.aws_region.current.name
          "awslogs-stream-prefix" = "nginx"
        }
      }
    }
  ]

  tags = [{
    key   = "Modified By"
    value = "AWSCC"
  }]
}