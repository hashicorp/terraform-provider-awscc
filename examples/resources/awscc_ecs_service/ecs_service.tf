data "awscc_ec2_subnet" "subnet" {
  id = "subnet-0000000"
}

resource "awscc_iam_role" "nginx" {
  role_name                   = "ecs_Task_ExecutionRole"
  assume_role_policy_document = <<EOF
{
    "Version": "2012-10-17",
    "Statement": [
      {
        "Action": "sts:AssumeRole",
        "Principal": {
          "Service": "ecs-tasks.amazonaws.com"
        },
        "Effect": "Allow",
        "Sid": ""
      }
    ]
}
  EOF
  managed_policy_arns         = ["arn:aws:iam::aws:policy/service-role/AmazonECSTaskExecutionRolePolicy"]
}

resource "awscc_ecs_service" "nginx" {
  service_name    = "nginx_service"
  cluster         = awscc_ecs_cluster.this.id
  task_definition = aws_ecs_task_definition.nginx.arn
  launch_type     = "FARGATE"
  desired_count   = 1

  network_configuration = {
    awsvpc_configuration = {
      assign_public_ip = "ENABLED"
      subnets          = ["${data.awscc_ec2_subnet.subnet.subnet_id}"]
    }
  }

  tags = [{
    key   = "Modified By"
    value = "AWSCC"
  }]
  depends_on = [awscc_iam_role.nginx]
}


