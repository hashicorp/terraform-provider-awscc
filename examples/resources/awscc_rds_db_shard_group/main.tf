data "aws_region" "current" {}

# Create a VPC
resource "awscc_ec2_vpc" "example" {
  cidr_block           = "10.0.0.0/16"
  enable_dns_hostnames = true
  enable_dns_support   = true
  tags = [{
    key   = "Modified By"
    value = "AWSCC"
  }]
}

# Create subnets
resource "awscc_ec2_subnet" "example1" {
  vpc_id            = awscc_ec2_vpc.example.id
  cidr_block        = "10.0.1.0/24"
  availability_zone = "${data.aws_region.current.name}a"
  tags = [{
    key   = "Modified By"
    value = "AWSCC"
  }]
}

resource "awscc_ec2_subnet" "example2" {
  vpc_id            = awscc_ec2_vpc.example.id
  cidr_block        = "10.0.2.0/24"
  availability_zone = "${data.aws_region.current.name}b"
  tags = [{
    key   = "Modified By"
    value = "AWSCC"
  }]
}

# Create DB subnet group
resource "awscc_rds_db_subnet_group" "example" {
  db_subnet_group_description = "Subnet group for Aurora Limitless"
  subnet_ids                  = [awscc_ec2_subnet.example1.id, awscc_ec2_subnet.example2.id]
  db_subnet_group_name        = "example-subnet-group"
  tags = [{
    key   = "Modified By"
    value = "AWSCC"
  }]
}

# Create an Aurora DB cluster
resource "awscc_rds_db_cluster" "example" {
  engine                 = "aurora-postgresql"
  engine_version         = "15.3"
  master_username        = "exampleuser"
  master_user_password   = "examplepassword123!"
  db_cluster_identifier  = "example-aurora-cluster"
  vpc_security_group_ids = [awscc_ec2_security_group.example.id]
  db_subnet_group_name   = awscc_rds_db_subnet_group.example.db_subnet_group_name
  tags = [{
    key   = "Modified By"
    value = "AWSCC"
  }]
}

# Create security group for the DB cluster
resource "awscc_ec2_security_group" "example" {
  group_description = "Security group for Aurora Limitless"
  vpc_id            = awscc_ec2_vpc.example.id
  security_group_ingress = [
    {
      from_port   = 5432
      to_port     = 5432
      ip_protocol = "tcp"
      cidr_ip     = "0.0.0.0/0"
    }
  ]
  tags = [{
    key   = "Modified By"
    value = "AWSCC"
  }]
}

# Create DB shard group
resource "awscc_rds_db_shard_group" "example" {
  db_cluster_identifier     = awscc_rds_db_cluster.example.id
  db_shard_group_identifier = "example-shard-group"
  max_acu                   = 128
  min_acu                   = 64
  compute_redundancy        = 2
  publicly_accessible       = false
  tags = [{
    key   = "Modified By"
    value = "AWSCC"
  }]
}