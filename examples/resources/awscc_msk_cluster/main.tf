data "aws_region" "current" {}

# VPC and Network Resources
resource "awscc_ec2_vpc" "msk_vpc" {
  cidr_block = "10.0.0.0/16"
  tags = [{
    key   = "Name"
    value = "msk-vpc"
  }]
}

resource "awscc_ec2_subnet" "msk_subnet_1" {
  vpc_id            = awscc_ec2_vpc.msk_vpc.id
  cidr_block        = "10.0.1.0/24"
  availability_zone = "${data.aws_region.current.name}a"
  tags = [{
    key   = "Name"
    value = "msk-subnet-1"
  }]
}

resource "awscc_ec2_subnet" "msk_subnet_2" {
  vpc_id            = awscc_ec2_vpc.msk_vpc.id
  cidr_block        = "10.0.2.0/24"
  availability_zone = "${data.aws_region.current.name}b"
  tags = [{
    key   = "Name"
    value = "msk-subnet-2"
  }]
}

# Security Group for MSK
resource "awscc_ec2_security_group" "msk_sg" {
  group_description = "Security group for MSK cluster"
  vpc_id            = awscc_ec2_vpc.msk_vpc.id
  security_group_ingress = [
    {
      ip_protocol = "tcp"
      from_port   = 9092
      to_port     = 9092
      cidr_ip     = awscc_ec2_vpc.msk_vpc.cidr_block
    }
  ]
  tags = [{
    key   = "Name"
    value = "msk-security-group"
  }]
}

# MSK Cluster
resource "awscc_msk_cluster" "example" {
  cluster_name           = "example-msk-cluster"
  kafka_version          = "2.8.1"
  number_of_broker_nodes = 2

  broker_node_group_info = {
    instance_type = "kafka.t3.small"
    client_subnets = [
      awscc_ec2_subnet.msk_subnet_1.id,
      awscc_ec2_subnet.msk_subnet_2.id
    ]
    security_groups = [awscc_ec2_security_group.msk_sg.id]
    storage_info = {
      ebs_storage_info = {
        volume_size = 100
      }
    }
  }

  encryption_info = {
    encryption_in_transit = {
      client_broker = "TLS"
      in_cluster    = true
    }
  }

  enhanced_monitoring = "DEFAULT"

  logging_info = {
    broker_logs = {
      cloudwatch_logs = {
        enabled = true
      }
    }
  }

  tags = [{
    key   = "Modified By"
    value = "AWSCC"
  }]
}