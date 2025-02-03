# Get current AWS region
data "aws_region" "current" {}

# Get current AWS account ID
data "aws_caller_identity" "current" {}

# Example DataBrew Schedule
resource "awscc_databrew_schedule" "example" {
  name = "example-schedule"

  # Run daily at midnight UTC
  cron_expression = "cron(0 0 * * ? *)"


  # Tags
  tags = [
    {
      key   = "Environment"
      value = "Production"
    },
    {
      key   = "Modified By"
      value = "AWSCC"
    }
  ]
}