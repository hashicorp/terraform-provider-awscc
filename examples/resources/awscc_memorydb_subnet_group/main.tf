data "aws_region" "current" {}

# Create a VPC for MemoryDB
resource "awscc_ec2_vpc" "memorydb" {
  cidr_block = "10.0.0.0/16"
  tags = [{
    key   = "Name"
    value = "memorydb-vpc"
  }]
}

# Create internet gateway and attach it to VPC
resource "awscc_ec2_internet_gateway" "memorydb" {
  tags = [{
    key   = "Name"
    value = "memorydb-igw"
  }]
}

resource "awscc_ec2_vpc_gateway_attachment" "igw" {
  vpc_id              = awscc_ec2_vpc.memorydb.id
  internet_gateway_id = awscc_ec2_internet_gateway.memorydb.id
}

# Create two subnets for MemoryDB
resource "awscc_ec2_subnet" "memorydb_1" {
  vpc_id            = awscc_ec2_vpc.memorydb.id
  cidr_block        = "10.0.1.0/24"
  availability_zone = "${data.aws_region.current.name}a"
  tags = [{
    key   = "Name"
    value = "memorydb-subnet-1"
  }]
}

resource "awscc_ec2_subnet" "memorydb_2" {
  vpc_id            = awscc_ec2_vpc.memorydb.id
  cidr_block        = "10.0.2.0/24"
  availability_zone = "${data.aws_region.current.name}b"
  tags = [{
    key   = "Name"
    value = "memorydb-subnet-2"
  }]
}

# Create MemoryDB subnet group
resource "awscc_memorydb_subnet_group" "example" {
  subnet_group_name = "example-memorydb-subnet-group"
  subnet_ids        = [awscc_ec2_subnet.memorydb_1.id, awscc_ec2_subnet.memorydb_2.id]
  description       = "Example MemoryDB subnet group"

  tags = [{
    key   = "Environment"
    value = "example"
  }]
}