resource "awscc_memorydb_user" "example" {
  user_name     = "sample-user"
  access_string = "on ~* +@all" # Full access to all commands and keys
  authentication_mode = {
    type      = "password"
    passwords = ["MyP@ssw0rd123!456789"] # You should use sensitive values in production
  }

  tags = [{
    key   = "Modified By"
    value = "AWSCC"
  }]
}