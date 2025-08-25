# VPC for our network resources
resource "aws_vpc" "example" {
  cidr_block = "10.0.0.0/16"

  tags = {
    Name        = "example-vpc"
    Environment = "example"
  }
}

# Two subnets to test connectivity between
resource "aws_subnet" "example_source" {
  vpc_id            = aws_vpc.example.id
  cidr_block        = "10.0.1.0/24"
  availability_zone = "us-west-2a"

  tags = {
    Name        = "example-source"
    Environment = "example"
  }
}

resource "aws_subnet" "example_destination" {
  vpc_id            = aws_vpc.example.id
  cidr_block        = "10.0.2.0/24"
  availability_zone = "us-west-2b"

  tags = {
    Name        = "example-destination"
    Environment = "example"
  }
}

# Security Group for instances
resource "aws_security_group" "example" {
  name        = "example-sg"
  description = "Allow traffic for network insights"
  vpc_id      = aws_vpc.example.id

  ingress {
    from_port   = 22
    to_port     = 22
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
    Name        = "example-sg"
    Environment = "example"
  }
}

# Source network interface
resource "aws_network_interface" "source" {
  subnet_id       = aws_subnet.example_source.id
  security_groups = [aws_security_group.example.id]

  tags = {
    Name        = "source-eni"
    Environment = "example"
  }
}

# Destination network interface
resource "aws_network_interface" "destination" {
  subnet_id       = aws_subnet.example_destination.id
  security_groups = [aws_security_group.example.id]

  tags = {
    Name        = "destination-eni"
    Environment = "example"
  }
}

# Network Insights Path - defines the connectivity path to analyze
resource "awscc_ec2_network_insights_path" "example" {
  source           = aws_network_interface.source.id
  destination      = aws_network_interface.destination.id
  protocol         = "tcp"
  destination_port = 22

  tags = [
    {
      key   = "Name"
      value = "example-path"
    },
    {
      key   = "Environment"
      value = "example"
    }
  ]
}

# Network Insights Analysis - performs analysis on the defined path
resource "awscc_ec2_network_insights_analysis" "example" {
  network_insights_path_id = awscc_ec2_network_insights_path.example.id

  tags = [
    {
      key   = "Name"
      value = "example-analysis"
    },
    {
      key   = "Environment"
      value = "example"
    }
  ]
}
