# Get current AWS account and region
data "aws_caller_identity" "current" {}
data "aws_region" "current" {}

# Create cross region replication configuration
resource "awscc_ecr_replication_configuration" "example" {
  replication_configuration = {
    rules = [
      {
        destinations = [
          {
            region      = (data.aws_region.current.name == "eu-west-1" ? "us-west-2" : "eu-west-1") # Primary replication region
            registry_id = data.aws_caller_identity.current.account_id
          },
          {
            region      = "ap-southeast-1" # Secondary replication region
            registry_id = data.aws_caller_identity.current.account_id
          }
        ]
        repository_filters = [
          {
            filter      = "prod-"
            filter_type = "PREFIX_MATCH"
          }
        ]
      }
    ]
  }
}