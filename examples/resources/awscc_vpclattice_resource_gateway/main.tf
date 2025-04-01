# Create a VPC for the resource gateway
resource "aws_vpc" "example" {
  cidr_block           = "10.0.0.0/16"
  enable_dns_hostnames = true
  enable_dns_support   = true

  tags = {
    Name = "vpc-lattice-example"
  }
}

# Create subnets for the resource gateway
resource "aws_subnet" "example" {
  count                   = 2
  vpc_id                  = aws_vpc.example.id
  cidr_block              = "10.0.${count.index + 1}.0/24"
  availability_zone       = data.aws_availability_zones.available.names[count.index]
  map_public_ip_on_launch = true

  tags = {
    Name = "subnet-lattice-example-${count.index + 1}"
  }
}

# Get available AZs
data "aws_availability_zones" "available" {
  state = "available"
}

# Create security group for the resource gateway
resource "aws_security_group" "example" {
  name_prefix = "vpclattice-rg-"
  vpc_id      = aws_vpc.example.id

  ingress {
    from_port   = 443
    to_port     = 443
    protocol    = "tcp"
    cidr_blocks = ["0.0.0.0/0"]
  }

  egress {
    from_port   = 0
    to_port     = 0
    protocol    = "-1"
    cidr_blocks = ["0.0.0.0/0"]
  }

  tags = {
    Name = "sg-lattice-example"
  }
}

# Create the VPC Lattice Resource Gateway
resource "awscc_vpclattice_resource_gateway" "example" {
  name               = "example-resource-gateway"
  vpc_identifier     = aws_vpc.example.id
  subnet_ids         = aws_subnet.example[*].id
  security_group_ids = [aws_security_group.example.id]
  ip_address_type    = "IPV4"

  tags = [{
    key   = "Modified By"
    value = "AWSCC"
  }]
}