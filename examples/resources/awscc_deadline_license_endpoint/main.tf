data "aws_region" "current" {}

# Create VPC
resource "awscc_ec2_vpc" "deadline" {
  cidr_block = "10.0.0.0/16"
  tags = [{
    key   = "Modified By"
    value = "AWSCC"
  }]
}

# Create Internet Gateway
resource "awscc_ec2_internet_gateway" "deadline" {
  tags = [{
    key   = "Modified By"
    value = "AWSCC"
  }]
}

# Attach Internet Gateway to VPC
resource "awscc_ec2_vpc_gateway_attachment" "deadline" {
  vpc_id              = awscc_ec2_vpc.deadline.id
  internet_gateway_id = awscc_ec2_internet_gateway.deadline.id
}

# Create subnets
resource "awscc_ec2_subnet" "deadline_1" {
  vpc_id            = awscc_ec2_vpc.deadline.id
  cidr_block        = "10.0.1.0/24"
  availability_zone = "${data.aws_region.current.name}a"
  tags = [{
    key   = "Modified By"
    value = "AWSCC"
  }]
}

resource "awscc_ec2_subnet" "deadline_2" {
  vpc_id            = awscc_ec2_vpc.deadline.id
  cidr_block        = "10.0.2.0/24"
  availability_zone = "${data.aws_region.current.name}b"
  tags = [{
    key   = "Modified By"
    value = "AWSCC"
  }]
}

# Create security group
resource "awscc_ec2_security_group" "deadline" {
  group_description = "Security group for Deadline License Endpoint"
  vpc_id            = awscc_ec2_vpc.deadline.id
  group_name        = "deadline-license-endpoint-sg"
  security_group_ingress = [{
    ip_protocol = "tcp"
    from_port   = 7246
    to_port     = 7246
    cidr_ip     = "10.0.0.0/16"
  }]
  tags = [{
    key   = "Modified By"
    value = "AWSCC"
  }]
}

# Create Deadline License Endpoint
resource "awscc_deadline_license_endpoint" "example" {
  vpc_id = awscc_ec2_vpc.deadline.id
  subnet_ids = [
    awscc_ec2_subnet.deadline_1.id,
    awscc_ec2_subnet.deadline_2.id
  ]
  security_group_ids = [awscc_ec2_security_group.deadline.id]
  tags = [{
    key   = "Modified By"
    value = "AWSCC"
  }]
}