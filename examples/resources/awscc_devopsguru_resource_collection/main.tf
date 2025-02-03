# Get current AWS account ID and region
data "aws_caller_identity" "current" {}
data "aws_region" "current" {}

# Create DevOps Guru Resource Collection
resource "awscc_devopsguru_resource_collection" "example" {
  resource_collection_filter = {
    cloudformation = {
      stack_names = ["example-stack"]
    }
    tags = [
      {
        app_boundary = {
          key   = "Environment"
          value = "Production"
        }
      }
    ]
  }
}