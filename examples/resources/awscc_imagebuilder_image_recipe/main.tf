# Data sources
data "aws_region" "current" {}

# Create an Image Recipe
resource "awscc_imagebuilder_image_recipe" "example" {
  name        = "example-recipe"
  description = "An example Image Recipe"
  version     = "1.0.0"

  # Using Amazon Linux 2023 as parent image
  parent_image = "arn:aws:imagebuilder:${data.aws_region.current.name}:aws:image/amazon-linux-2023-x86/x.x.x"

  # Components that will be part of the image
  components = [
    {
      component_arn = "arn:aws:imagebuilder:${data.aws_region.current.name}:aws:component/amazon-cloudwatch-agent-linux/x.x.x"
    },
    {
      component_arn = "arn:aws:imagebuilder:${data.aws_region.current.name}:aws:component/aws-cli-version-2-linux/x.x.x"
    }
  ]

  # Configure block device mappings
  block_device_mappings = [
    {
      device_name = "/dev/xvda"
      ebs = {
        volume_size = 30
        volume_type = "gp3"
        encrypted   = true
      }
    }
  ]

  # Additional instance configuration
  additional_instance_configuration = {
    systems_manager_agent = {
      uninstall_after_build = false
    }
  }

  # Working directory for build and test workflows
  working_directory = "/tmp"

  # Tags
  tags = [{
    key   = "Modified By"
    value = "AWSCC"
  }]
}