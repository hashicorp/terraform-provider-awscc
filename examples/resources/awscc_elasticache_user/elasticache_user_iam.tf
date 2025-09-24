resource "awscc_elasticache_user" "example" {
  user_id       = "testuserid"
  user_name     = "testuserid"
  access_string = "on ~* +@all"
  engine        = "redis"

  authentication_mode = {
    type = "iam"
  }

  tags = [{
    key   = "Modified By"
    value = "AWSCC"
  }]
}