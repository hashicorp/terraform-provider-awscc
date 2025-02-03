# Get current AWS region and account
data "aws_region" "current" {}
data "aws_caller_identity" "current" {}

# Create an EC2 Image Builder distribution configuration
resource "awscc_imagebuilder_distribution_configuration" "example" {
  name        = "example-dist-config"
  description = "Example distribution configuration for Image Builder"

  distributions = [
    {
      region = data.aws_region.current.name
      ami_distribution_configuration = {
        name               = "example-ami-{{ imagebuilder:buildDate }}"
        description        = "AMI created by Image Builder"
        target_account_ids = [data.aws_caller_identity.current.account_id]

        launch_permission_configuration = {
          user_ids = [data.aws_caller_identity.current.account_id]
        }

        ami_tags = {
          "Environment" = "Test"
          "Created-By"  = "AWSCC"
        }
      }
    }
  ]

  tags = [{
    key   = "Environment"
    value = "Test"
  }]
}