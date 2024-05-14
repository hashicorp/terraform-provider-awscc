resource "awscc_elasticache_user" "example" {
  user_id       = "testuserid"
  user_name     = "testUserName"
  access_string = "on ~app::* -@all +@read +@hash +@bitmap +@geo -setbit -bitfield -hset -hsetnx -hmset -hincrby -hincrbyfloat -hdel -bitop -geoadd -georadius -georadiusbymember"
  engine        = "redis"
  passwords     = ["password123456789"]

  tags = [{
    key   = "Modified By"
    value = "AWSCC"
  }]
}