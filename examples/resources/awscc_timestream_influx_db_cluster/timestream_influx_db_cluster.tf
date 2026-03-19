# VPC for InfluxDB cluster
resource "awscc_ec2_vpc" "example" {
  cidr_block           = "10.0.0.0/16"
  enable_dns_hostnames = true
  enable_dns_support   = true
  tags = [
    {
      key   = "Name"
      value = "example-influx-vpc"
    },
    {
      key   = "Environment"
      value = "example"
    }
  ]
}

# Subnet 1 for InfluxDB cluster
resource "awscc_ec2_subnet" "example_1" {
  vpc_id            = awscc_ec2_vpc.example.id
  cidr_block        = "10.0.1.0/24"
  availability_zone = "us-west-2a"
  tags = [
    {
      key   = "Name"
      value = "example-influx-subnet-1"
    },
    {
      key   = "Environment"
      value = "example"
    }
  ]
}

# Subnet 2 for InfluxDB cluster
resource "awscc_ec2_subnet" "example_2" {
  vpc_id            = awscc_ec2_vpc.example.id
  cidr_block        = "10.0.2.0/24"
  availability_zone = "us-west-2b"
  tags = [
    {
      key   = "Name"
      value = "example-influx-subnet-2"
    },
    {
      key   = "Environment"
      value = "example"
    }
  ]
}

# Security group for InfluxDB cluster
resource "awscc_ec2_security_group" "example" {
  group_description = "Security group for InfluxDB cluster"
  vpc_id            = awscc_ec2_vpc.example.id
  tags = [
    {
      key   = "Name"
      value = "example-influx-sg"
    },
    {
      key   = "Environment"
      value = "example"
    }
  ]
}

# Security group ingress rule for InfluxDB port
resource "awscc_ec2_security_group_ingress" "influx_port" {
  group_id    = awscc_ec2_security_group.example.id
  cidr_ip     = awscc_ec2_vpc.example.cidr_block
  from_port   = 8086
  to_port     = 8086
  ip_protocol = "tcp"
}

# Security group egress rule for all outbound traffic
resource "awscc_ec2_security_group_egress" "all_outbound" {
  group_id    = awscc_ec2_security_group.example.id
  cidr_ip     = "0.0.0.0/0"
  ip_protocol = "-1"
}

# InfluxDB Cluster
resource "awscc_timestream_influx_db_cluster" "example" {
  name                = "example-influx-cluster"
  username            = "admin"
  password            = "ExamplePassword123"
  organization        = "example-org"
  bucket              = "example-bucket"
  db_instance_type    = "db.influx.medium"
  db_storage_type     = "InfluxIOIncludedT1"
  allocated_storage   = 200
  deployment_type     = "MULTI_NODE_READ_REPLICAS"
  publicly_accessible = false
  network_type        = "IPV4"

  vpc_subnet_ids = [
    awscc_ec2_subnet.example_1.id,
    awscc_ec2_subnet.example_2.id
  ]

  vpc_security_group_ids = [
    awscc_ec2_security_group.example.id
  ]

  tags = [
    {
      key   = "Name"
      value = "example-influx-cluster"
    },
    {
      key   = "Environment"
      value = "example"
    }
  ]
}
