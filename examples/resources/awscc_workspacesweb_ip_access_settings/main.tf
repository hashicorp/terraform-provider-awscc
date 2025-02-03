# Get current region
data "aws_region" "current" {}

# Get current account ID
data "aws_caller_identity" "current" {}

# Create the IP access settings
resource "awscc_workspacesweb_ip_access_settings" "example" {
  display_name = "example-ip-access-settings"
  description  = "Example IP access settings for WorkSpaces Web"

  ip_rules = [
    {
      ip_range    = "192.168.0.0/24"
      description = "Allow internal network access"
    },
    {
      ip_range    = "10.0.0.0/8"
      description = "Allow VPC network access"
    }
  ]

  tags = [{
    key   = "Modified By"
    value = "AWSCC"
  }]
}

# Output the IP access settings ID and ARN
output "ip_access_settings_id" {
  value = awscc_workspacesweb_ip_access_settings.example.id
}

output "ip_access_settings_arn" {
  value = awscc_workspacesweb_ip_access_settings.example.ip_access_settings_arn
}