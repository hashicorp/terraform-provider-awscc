# Create CodeDeploy Application
resource "awscc_codedeploy_application" "example" {
  application_name = "example-application"
  compute_platform = "Server" # Valid values: Server, Lambda, or ECS

  tags = [{
    key   = "Modified By"
    value = "AWSCC"
  }]
}