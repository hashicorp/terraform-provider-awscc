# VPC and Networking Resources
resource "aws_vpc" "main" {
  cidr_block = "10.0.0.0/16"

  tags = {
    Name = "ECS Task Set Demo VPC"
  }
}

resource "aws_subnet" "public1" {
  vpc_id            = aws_vpc.main.id
  cidr_block        = "10.0.1.0/24"
  availability_zone = "us-west-2a"

  tags = {
    Name = "ECS Task Set Public Subnet 1"
  }
}

resource "aws_subnet" "public2" {
  vpc_id            = aws_vpc.main.id
  cidr_block        = "10.0.2.0/24"
  availability_zone = "us-west-2b"

  tags = {
    Name = "ECS Task Set Public Subnet 2"
  }
}

resource "aws_internet_gateway" "main" {
  vpc_id = aws_vpc.main.id

  tags = {
    Name = "ECS Task Set IGW"
  }
}

resource "aws_security_group" "ecs" {
  name        = "ecs-task-set-sg"
  description = "Security group for ECS tasks"
  vpc_id      = aws_vpc.main.id

  ingress {
    from_port   = 80
    to_port     = 80
    protocol    = "tcp"
    cidr_blocks = ["0.0.0.0/0"]
  }

  egress {
    from_port   = 0
    to_port     = 0
    protocol    = "-1"
    cidr_blocks = ["0.0.0.0/0"]
  }
}

# Load Balancer Resources
resource "aws_lb" "main" {
  name               = "ecs-task-set-alb"
  internal           = false
  load_balancer_type = "application"
  security_groups    = [aws_security_group.ecs.id]
  subnets            = [aws_subnet.public1.id, aws_subnet.public2.id]
}

resource "aws_lb_target_group" "main" {
  name        = "ecs-task-set-tg"
  port        = 80
  protocol    = "HTTP"
  vpc_id      = aws_vpc.main.id
  target_type = "ip"
}

resource "aws_lb_listener" "main" {
  load_balancer_arn = aws_lb.main.arn
  port              = "80"
  protocol          = "HTTP"

  default_action {
    type             = "forward"
    target_group_arn = aws_lb_target_group.main.arn
  }
}

# IAM Resources
resource "awscc_iam_role" "ecs_task_execution" {
  role_name = "ecs-task-set-execution-role"
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
  managed_policy_arns = ["arn:aws:iam::aws:policy/service-role/AmazonECSTaskExecutionRolePolicy"]

  tags = [{
    key   = "Modified By"
    value = "AWSCC"
  }]
}

resource "awscc_iam_role" "ecs_task" {
  role_name = "ecs-task-set-role"
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

  tags = [{
    key   = "Modified By"
    value = "AWSCC"
  }]
}

# ECS Resources
resource "awscc_ecs_cluster" "main" {
  cluster_name = "task-set-demo"

  tags = [{
    key   = "Modified By"
    value = "AWSCC"
  }]
}

resource "awscc_ecs_task_definition" "app" {
  family = "task-set-app"

  container_definitions = [
    {
      name      = "nginx"
      image     = "nginx:latest"
      cpu       = 256
      memory    = 512
      essential = true
      port_mappings = [
        {
          container_port = 80
          protocol       = "tcp"
        }
      ]
    }
  ]

  requires_compatibilities = ["FARGATE"]
  network_mode             = "awsvpc"
  cpu                      = "512"
  memory                   = "1024"
  execution_role_arn       = awscc_iam_role.ecs_task_execution.arn
  task_role_arn            = awscc_iam_role.ecs_task.arn

  tags = [{
    key   = "Modified By"
    value = "AWSCC"
  }]
}

resource "awscc_ecs_service" "app" {
  service_name  = "task-set-service"
  cluster       = awscc_ecs_cluster.main.id
  desired_count = 2

  deployment_configuration = {
    deployment_circuit_breaker = {
      enable   = true
      rollback = true
    }
    maximum_percent         = 200
    minimum_healthy_percent = 100
  }

  load_balancers = [
    {
      container_name   = "nginx"
      container_port   = 80
      target_group_arn = aws_lb_target_group.main.arn
    }
  ]

  network_configuration = {
    aws_vpc_configuration = {
      assign_public_ip = true
      security_groups  = [aws_security_group.ecs.id]
      subnets          = [aws_subnet.public1.id, aws_subnet.public2.id]
    }
  }

  task_definition = awscc_ecs_task_definition.app.id

  scheduling_strategy = "REPLICA"

  tags = [{
    key   = "Modified By"
    value = "AWSCC"
  }]
}

resource "awscc_ecs_task_set" "app_v1" {
  service     = awscc_ecs_service.app.id
  cluster     = awscc_ecs_cluster.main.id
  external_id = "v1"
  scale = {
    unit  = "PERCENT"
    value = 100
  }
  task_definition = awscc_ecs_task_definition.app.id

  tags = [{
    key   = "Modified By"
    value = "AWSCC"
  }]
}