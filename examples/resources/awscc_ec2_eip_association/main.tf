data "aws_region" "current" {}

# Create a VPC
resource "awscc_ec2_vpc" "example" {
  cidr_block = "10.0.0.0/16"
  tags = [{
    key   = "Modified By"
    value = "AWSCC"
  }]
}

# Create an Internet Gateway
resource "awscc_ec2_internet_gateway" "example" {
  tags = [{
    key   = "Modified By"
    value = "AWSCC"
  }]
}

# Attach Internet Gateway to VPC
resource "awscc_ec2_vpc_gateway_attachment" "example" {
  vpc_id              = awscc_ec2_vpc.example.id
  internet_gateway_id = awscc_ec2_internet_gateway.example.id
}

# Create a subnet
resource "awscc_ec2_subnet" "example" {
  vpc_id            = awscc_ec2_vpc.example.id
  cidr_block        = "10.0.1.0/24"
  availability_zone = "${data.aws_region.current.name}a"
  tags = [{
    key   = "Modified By"
    value = "AWSCC"
  }]
}

# Create a route table
resource "awscc_ec2_route_table" "example" {
  vpc_id = awscc_ec2_vpc.example.id
  tags = [{
    key   = "Modified By"
    value = "AWSCC"
  }]
}

# Create a route to the internet gateway
resource "awscc_ec2_route" "internet_gateway" {
  route_table_id         = awscc_ec2_route_table.example.id
  destination_cidr_block = "0.0.0.0/0"
  gateway_id             = awscc_ec2_internet_gateway.example.id
}

# Associate the route table with the subnet
resource "awscc_ec2_subnet_route_table_association" "example" {
  route_table_id = awscc_ec2_route_table.example.id
  subnet_id      = awscc_ec2_subnet.example.id
}

# Create a security group
resource "awscc_ec2_security_group" "example" {
  group_description = "Example security group"
  vpc_id            = awscc_ec2_vpc.example.id
  group_name        = "example-sg"
  security_group_ingress = [
    {
      ip_protocol = -1
      from_port   = -1
      to_port     = -1
      cidr_ip     = "0.0.0.0/0"
    }
  ]
  tags = [{
    key   = "Modified By"
    value = "AWSCC"
  }]
}

# Create an EC2 instance
resource "awscc_ec2_instance" "example" {
  instance_type = "t2.micro"
  subnet_id     = awscc_ec2_subnet.example.id
  image_id      = "ami-0735c191cf914754d"
  tags = [{
    key   = "Modified By"
    value = "AWSCC"
  }]
  security_group_ids = [awscc_ec2_security_group.example.id]
}

# Create an Elastic IP
resource "awscc_ec2_eip" "example" {
  domain = "vpc"
  tags = [{
    key   = "Modified By"
    value = "AWSCC"
  }]
}

# Create EIP association
resource "awscc_ec2_eip_association" "example" {
  allocation_id = awscc_ec2_eip.example.allocation_id
  instance_id   = awscc_ec2_instance.example.id
}