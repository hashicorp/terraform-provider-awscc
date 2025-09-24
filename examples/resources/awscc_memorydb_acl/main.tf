# Create a MemoryDB user
resource "awscc_memorydb_user" "example" {
  user_name     = "example-user"
  access_string = "on ~* +@all"
  authentication_mode = {
    type      = "password"
    passwords = ["MySecurePassword123!!@@"]
  }
  tags = [{
    key   = "Modified By"
    value = "AWSCC"
  }]
}

# Create a MemoryDB ACL
resource "awscc_memorydb_acl" "example" {
  acl_name = "example-acl"
  tags = [{
    key   = "Environment"
    value = "Test"
    },
    {
      key   = "Modified By"
      value = "AWSCC"
  }]
  user_names = [awscc_memorydb_user.example.user_name]
}