data "aws_vpc" "default" {
  default = true
}

data "aws_security_group" "default" {
  name   = "default"
  vpc_id = data.aws_vpc.default.id
}

resource "awscc_rds_option_group" "example_rds_option_group" {
  engine_name              = "mysql"
  major_engine_version     = "8.0"
  option_group_description = "Example MySQL RDS option group using Memcached"
  option_configurations = [{
    option_name                    = "MEMCACHED"
    vpc_security_group_memberships = [data.aws_security_group.default.id]
    port                           = 3306
    option_settings = [{
      name  = "CHUNK_SIZE"
      value = "32"
      },
      {
        name  = "BINDING_PROTOCOL"
        value = "ascii"
    }]
  }]
  tags = [{
    key   = "Modified By"
    value = "AWSCC"
  }]

}