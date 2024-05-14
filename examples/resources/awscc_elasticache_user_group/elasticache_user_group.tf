resource "awscc_elasticache_user_group" "example" {
  engine        = "redis"
  user_group_id = "example-group"
  user_ids      = [awscc_elasticache_user.example1.user_id, awscc_elasticache_user.example2.user_id]

  tags = [{
    key   = "Modified By"
    value = "AWSCC"
  }]

}


resource "awscc_elasticache_user" "example1" {
  user_id       = "example"
  user_name     = "example"
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

resource "awscc_elasticache_user" "example2" {
  user_id       = "example2"
  user_name     = "default"
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