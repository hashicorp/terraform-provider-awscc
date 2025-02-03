# Get current AWS region
data "aws_region" "current" {}

# Get current AWS account ID
data "aws_caller_identity" "current" {}

# Example Image Builder workflow with minimal required attributes
resource "awscc_imagebuilder_workflow" "example" {
  name    = "example-workflow"
  type    = "TEST" # Can be BUILD, TEST, or DISTRIBUTE
  version = "1.0.0"

  data = jsonencode({
    name          = "example-workflow"
    schemaVersion = 1.0
    steps = [
      {
        name   = "test-step"
        action = "WaitForAction"
      }
    ]
  })

  tags = {
    Environment = "test"
    Modified_By = "AWSCC"
  }
}