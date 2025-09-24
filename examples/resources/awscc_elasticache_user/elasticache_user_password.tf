resource "awscc_elasticache_user" "example" {
  user_id       = "testuserid"
  user_name     = "testUserName"
  access_string = "on ~* +@all"
  engine        = "redis"

  authentication_mode = {
    type      = "password"
    passwords = ["password123456789", "password987654321"]
  }

  tags = [{
    key   = "Modified By"
    value = "AWSCC"
  }]
}