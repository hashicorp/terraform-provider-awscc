data "aws_region" "current" {}

# VPC and Network Configuration
resource "awscc_ec2_vpc" "m2_vpc" {
  cidr_block           = "10.0.0.0/16"
  enable_dns_hostnames = true
  enable_dns_support   = true
  tags = [{
    key   = "Name"
    value = "m2-environment-vpc"
  }]
}

resource "awscc_ec2_internet_gateway" "m2_igw" {
  tags = [{
    key   = "Name"
    value = "m2-environment-igw"
  }]
}

resource "awscc_ec2_vpc_gateway_attachment" "m2_igw_attachment" {
  internet_gateway_id = awscc_ec2_internet_gateway.m2_igw.id
  vpc_id             = awscc_ec2_vpc.m2_vpc.id
}

resource "awscc_ec2_subnet" "m2_subnet_1" {
  vpc_id            = awscc_ec2_vpc.m2_vpc.id
  cidr_block        = "10.0.1.0/24"
  availability_zone = "${data.aws_region.current.name}a"
  tags = [{
    key   = "Name"
    value = "m2-environment-subnet-1"
  }]
}

resource "awscc_ec2_subnet" "m2_subnet_2" {
  vpc_id            = awscc_ec2_vpc.m2_vpc.id
  cidr_block        = "10.0.2.0/24"
  availability_zone = "${data.aws_region.current.name}b"
  tags = [{
    key   = "Name"
    value = "m2-environment-subnet-2"
  }]
}

# Security Group
resource "awscc_ec2_security_group" "m2_sg" {
  group_description = "Security group for M2 Environment"
  vpc_id           = awscc_ec2_vpc.m2_vpc.id
  security_group_egress = [{
    ip_protocol = "-1"
    from_port   = 0
    to_port     = 0
    cidr_ip     = "0.0.0.0/0"
  }]
  group_name = "m2-environment-sg"
  tags = [{
    key   = "Name"
    value = "m2-environment-sg"
  }]
}

# KMS Key
resource "awscc_kms_key" "m2_key" {
  description            = "KMS key for M2 Environment"
  pending_window_in_days = 7
  tags = [{
    key   = "Name"
    value = "m2-environment-key"
  }]
}

# M2 Environment
resource "awscc_m2_environment" "example" {
  name          = "my-m2-environment"
  engine_type   = "microfocus"
  instance_type = "m5.xlarge"
  description   = "Example M2 Environment for mainframe applications"

  high_availability_config = {
    desired_capacity = 2
  }

  subnet_ids         = [awscc_ec2_subnet.m2_subnet_1.id, awscc_ec2_subnet.m2_subnet_2.id]
  security_group_ids = [awscc_ec2_security_group.m2_sg.id]
  kms_key_id        = awscc_kms_key.m2_key.arn

  publicly_accessible = false
  network_type       = "dual"

  preferred_maintenance_window = "sun:23:00-mon:01:30"

  tags = [{
    key   = "Environment"
    value = "test"
  }, {
    key   = "Modified_By"
    value = "AWSCC"
  }]
}