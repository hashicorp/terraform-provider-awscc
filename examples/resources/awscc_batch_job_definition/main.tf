# Get current region name
data "aws_region" "current" {}

# Create an IAM role for the batch job
data "awscc_iam_policy_document" "batch_assume_role" {
  statement {
    effect = "Allow"
    principals {
      type        = "Service"
      identifiers = ["ecs-tasks.amazonaws.com"]
    }
    actions = ["sts:AssumeRole"]
  }
}

resource "awscc_iam_role" "batch_job_role" {
  role_name                   = "batchjob-example-role"
  assume_role_policy_document = data.awscc_iam_policy_document.batch_assume_role.json
  managed_policy_arns         = ["arn:aws:iam::aws:policy/service-role/AmazonECSTaskExecutionRolePolicy"]
  tags = [{
    key   = "Modified By"
    value = "AWSCC"
  }]
}

# Create the batch job definition
resource "awscc_batch_job_definition" "example" {
  job_definition_name = "example-job-definition"
  type                = "container"

  container_properties = jsonencode({
    command = ["echo", "Hello, World!"]
    image   = "amazonlinux:2"
    environment = [
      {
        name  = "REGION"
        value = data.aws_region.current.name
      }
    ]
    executionRoleArn = awscc_iam_role.batch_job_role.arn
    jobRoleArn       = awscc_iam_role.batch_job_role.arn
    memory           = 2048
    vcpus            = 1
  })

  retry_strategy {
    attempts = 2
  }

  timeout {
    attempt_duration_seconds = 300
  }

  tags = [{
    key   = "Modified By"
    value = "AWSCC"
  }]
}