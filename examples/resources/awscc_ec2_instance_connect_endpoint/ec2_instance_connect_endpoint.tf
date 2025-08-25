# AWS EC2 Instance Connect Endpoint Configuration

# VPC and subnet for the example
resource "aws_vpc" "example" {
  cidr_block = "10.0.0.0/16"
  tags = {
    Name        = "example-vpc"
    Environment = "example"
  }
}

resource "aws_subnet" "example" {
  vpc_id            = aws_vpc.example.id
  cidr_block        = "10.0.1.0/24"
  availability_zone = "us-west-2a"
  tags = {
    Name        = "example-subnet"
    Environment = "example"
  }
}

# Security group for the EC2 Instance Connect Endpoint
resource "aws_security_group" "example" {
  name        = "example-eice-sg"
  description = "Security group for EC2 Instance Connect Endpoint"
  vpc_id      = aws_vpc.example.id

  # Allow SSH traffic from the endpoint
  ingress {
    from_port   = 22
    to_port     = 22
    protocol    = "tcp"
    cidr_blocks = ["0.0.0.0/0"]
    description = "Allow SSH access"
  }

  egress {
    from_port   = 0
    to_port     = 0
    protocol    = "-1"
    cidr_blocks = ["0.0.0.0/0"]
    description = "Allow all outbound traffic"
  }

  tags = {
    Name        = "example-eice-sg"
    Environment = "example"
  }
}

# EC2 Instance Connect Endpoint
resource "awscc_ec2_instance_connect_endpoint" "example" {
  subnet_id          = aws_subnet.example.id
  security_group_ids = [aws_security_group.example.id]
  preserve_client_ip = false

  tags = [
    {
      key   = "Name"
      value = "example-endpoint"
    },
    {
      key   = "Environment"
      value = "example"
    }
  ]
}
