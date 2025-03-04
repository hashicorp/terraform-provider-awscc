data "aws_caller_identity" "current" {}
data "aws_region" "current" {}

# VPC and networking components
resource "awscc_ec2_vpc" "pcs" {
  cidr_block           = "10.0.0.0/16"
  enable_dns_hostnames = true
  enable_dns_support   = true

  tags = [{
    key   = "Name"
    value = "pcs-cluster-vpc"
  }]
}

resource "awscc_ec2_internet_gateway" "pcs" {
  tags = [{
    key   = "Name"
    value = "pcs-cluster-igw"
  }]
}

resource "awscc_ec2_vpc_gateway_attachment" "pcs" {
  vpc_id              = awscc_ec2_vpc.pcs.id
  internet_gateway_id = awscc_ec2_internet_gateway.pcs.id
}

resource "awscc_ec2_subnet" "pcs" {
  vpc_id                  = awscc_ec2_vpc.pcs.id
  cidr_block              = "10.0.1.0/24"
  availability_zone       = "${data.aws_region.current.name}a"
  map_public_ip_on_launch = true

  tags = [{
    key   = "Name"
    value = "pcs-cluster-subnet"
  }]
}

resource "awscc_ec2_route_table" "pcs" {
  vpc_id = awscc_ec2_vpc.pcs.id
  
  tags = [{
    key   = "Name"
    value = "pcs-cluster-rt"
  }]
}

resource "awscc_ec2_route" "internet" {
  route_table_id         = awscc_ec2_route_table.pcs.id
  destination_cidr_block = "0.0.0.0/0"
  gateway_id            = awscc_ec2_internet_gateway.pcs.id
}

resource "awscc_ec2_subnet_route_table_association" "pcs" {
  subnet_id      = awscc_ec2_subnet.pcs.id
  route_table_id = awscc_ec2_route_table.pcs.id
}

resource "awscc_ec2_security_group" "pcs" {
  group_description = "Security group for PCS cluster"
  vpc_id           = awscc_ec2_vpc.pcs.id

  security_group_ingress = [{
    ip_protocol = "tcp"
    from_port   = 22
    to_port     = 22
    cidr_ip     = awscc_ec2_vpc.pcs.cidr_block
  }]

  security_group_egress = [{
    ip_protocol = "-1"
    from_port   = -1
    to_port     = -1
    cidr_ip     = "0.0.0.0/0"
  }]

  tags = [{
    key   = "Name"
    value = "pcs-cluster-sg"
  }]
}

resource "awscc_pcs_cluster" "example" {
  name = "example-cluster"

  networking = {
    subnet_ids         = [awscc_ec2_subnet.pcs.id]
    security_group_ids = [awscc_ec2_security_group.pcs.id]
  }

  scheduler = {
    type    = "SLURM"
    version = "23.11"
  }

  size = "SMALL"

  tags = {
    Environment = "Test"
  }
}