# Create a VPC and subnet for our resources
resource "aws_vpc" "example" {
  cidr_block = "10.0.0.0/16"

  tags = {
    Name        = "example-vpc"
    Environment = "example"
  }
}

resource "aws_subnet" "example" {
  vpc_id     = aws_vpc.example.id
  cidr_block = "10.0.1.0/24"

  tags = {
    Name        = "example-subnet"
    Environment = "example"
  }
}

# Add an internet gateway for the EC2 instance
resource "aws_internet_gateway" "example" {
  vpc_id = aws_vpc.example.id

  tags = {
    Name        = "example-igw"
    Environment = "example"
  }
}

# Add a security group for the EC2 instance
resource "aws_security_group" "example" {
  name        = "example-sg"
  description = "Allow SSH and all outbound traffic"
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

# Create an EC2 instance
data "aws_ami" "amazon_linux" {
  most_recent = true
  owners      = ["amazon"]

  filter {
    name   = "name"
    values = ["amzn2-ami-hvm-*-x86_64-gp2"]
  }

  filter {
    name   = "virtualization-type"
    values = ["hvm"]
  }
}

resource "aws_instance" "example" {
  ami                    = data.aws_ami.amazon_linux.id
  instance_type          = "t2.micro"
  subnet_id              = aws_subnet.example.id
  vpc_security_group_ids = [aws_security_group.example.id]

  tags = {
    Name        = "example-instance"
    Environment = "example"
  }
}

# Create a network interface for traffic mirroring and attach it to the instance
resource "aws_network_interface" "example" {
  subnet_id       = aws_subnet.example.id
  description     = "Example network interface for traffic mirroring"
  security_groups = [aws_security_group.example.id]

  attachment {
    instance     = aws_instance.example.id
    device_index = 1
  }

  tags = {
    Name        = "example-mirror-interface"
    Environment = "example"
  }
}

# Network Load Balancer target configuration
resource "aws_lb" "example" {
  name               = "example-nlb"
  internal           = true
  load_balancer_type = "network"
  subnets            = [aws_subnet.example.id]

  tags = {
    Name        = "example-nlb"
    Environment = "example"
  }
}

resource "awscc_ec2_traffic_mirror_target" "eni_target" {
  description          = "Example Traffic Mirror Target using network interface"
  network_interface_id = aws_network_interface.example.id

  tags = [
    {
      key   = "Name"
      value = "example-eni-target"
    },
    {
      key   = "Environment"
      value = "example"
    }
  ]
}

resource "awscc_ec2_traffic_mirror_target" "nlb_target" {
  description               = "Example Traffic Mirror Target using NLB"
  network_load_balancer_arn = aws_lb.example.arn

  tags = [
    {
      key   = "Name"
      value = "example-nlb-target"
    },
    {
      key   = "Environment"
      value = "example"
    }
  ]
}
