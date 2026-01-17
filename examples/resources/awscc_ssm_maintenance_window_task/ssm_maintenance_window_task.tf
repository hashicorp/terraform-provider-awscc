# Create maintenance window first
resource "aws_ssm_maintenance_window" "example" {
  name     = "example-maintenance-window"
  schedule = "cron(0 2 ? * SUN *)"
  duration = 2
  cutoff   = 0

  tags = {
    Environment = "test"
  }
}

# Create maintenance window target
resource "aws_ssm_maintenance_window_target" "example" {
  window_id     = aws_ssm_maintenance_window.example.id
  resource_type = "INSTANCE"

  targets {
    key    = "tag:Environment"
    values = ["test"]
  }

  name        = "example-maintenance-target"
  description = "Example maintenance window target"
}

# Create maintenance window task
resource "awscc_ssm_maintenance_window_task" "example" {
  window_id       = aws_ssm_maintenance_window.example.id
  task_type       = "RUN_COMMAND"
  task_arn        = "AWS-RunShellScript"
  max_concurrency = "1"
  max_errors      = "0"

  targets = [{
    key    = "WindowTargetIds"
    values = [aws_ssm_maintenance_window_target.example.id]
  }]

  task_invocation_parameters = {
    maintenance_window_run_command_parameters = {
      parameters = jsonencode({
        commands = ["echo 'Hello World'"]
      })
    }
  }

  name        = "example-maintenance-task"
  description = "Example maintenance window task"
  priority    = 1
}
