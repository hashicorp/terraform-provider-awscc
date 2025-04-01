# Create a user settings configuration
resource "awscc_workspacesweb_user_settings" "example" {
  # Required settings
  copy_allowed     = "Enabled"
  download_allowed = "Enabled"
  paste_allowed    = "Enabled"
  print_allowed    = "Enabled"
  upload_allowed   = "Enabled"

  # Optional settings
  deep_link_allowed                  = "Enabled"
  disconnect_timeout_in_minutes      = 60
  idle_disconnect_timeout_in_minutes = 15

  # Cookie synchronization configuration example
  cookie_synchronization_configuration = {
    allowlist = [
      {
        domain = ".example.com"
        name   = "session"
        path   = "/"
      }
    ]
    blocklist = [
      {
        domain = ".blocked-example.com"
        name   = "tracking"
        path   = "/"
      }
    ]
  }

  # Tags
  tags = [{
    key   = "Modified By"
    value = "AWSCC"
  }]
}