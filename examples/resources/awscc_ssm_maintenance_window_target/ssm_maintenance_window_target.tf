# Create maintenance window first (using standard AWS provider)
resource "aws_ssm_maintenance_window" "example" {
  name     = "example-maintenance-window"
  schedule = "cron(0 2 ? * SUN *)"
  duration = 2
  cutoff   = 0
  
  tags = {
    Environment = "test"
  }
}

# Create maintenance window target (using AWSCC provider)
resource "awscc_ssm_maintenance_window_target" "example" {
  window_id     = aws_ssm_maintenance_window.example.id
  resource_type = "INSTANCE"
  
  targets = [{
    key    = "tag:Environment"
    values = ["test"]
  }]
  
  name        = "example-maintenance-target"
  description = "Example maintenance window target"
}
