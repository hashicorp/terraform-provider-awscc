# WorkspacesInstances Volume
resource "awscc_workspacesinstances_volume" "example" {
  availability_zone = "us-west-2a"
  size_in_gb        = 50
  volume_type       = "gp3"
  encrypted         = true

  # Optional performance settings
  iops = 3000

  tag_specifications = [
    {
      resource_type = "volume"
      tags = [
        {
          key   = "Name"
          value = "example-volume"
        },
        {
          key   = "Environment"
          value = "Test"
        }
      ]
    }
  ]
}
