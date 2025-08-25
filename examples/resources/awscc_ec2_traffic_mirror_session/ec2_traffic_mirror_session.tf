# Create a basic VPC
resource "aws_vpc" "example" {
  cidr_block = "10.0.0.0/16"
  tags = {
    Name = "example-vpc"
  }
}

# Create subnet for the NLB
resource "aws_subnet" "example" {
  vpc_id            = aws_vpc.example.id
  cidr_block        = "10.0.1.0/24"
  availability_zone = "us-west-2a"
  tags = {
    Name = "example-subnet"
  }
}

# Create a Network Load Balancer for use as a Traffic Mirror Target
resource "aws_lb" "example" {
  name               = "example-nlb"
  internal           = true
  load_balancer_type = "network"
  subnets            = [aws_subnet.example.id]

  tags = {
    Name = "example-nlb"
  }
}

# Create a security group for the EC2 instance
resource "aws_security_group" "example" {
  name        = "example-sg"
  description = "Allow basic traffic"
  vpc_id      = aws_vpc.example.id

  egress {
    from_port   = 0
    to_port     = 0
    protocol    = "-1"
    cidr_blocks = ["0.0.0.0/0"]
  }

  tags = {
    Name = "example-sg"
  }
}

# Use latest Amazon Linux 2 AMI
data "aws_ami" "amazon_linux" {
  most_recent = true
  owners      = ["amazon"]

  filter {
    name   = "name"
    values = ["amzn2-ami-hvm-*-x86_64-gp2"]
  }
}

# Create an EC2 instance
resource "aws_instance" "example" {
  ami             = data.aws_ami.amazon_linux.id
  instance_type   = "t3.micro"
  subnet_id       = aws_subnet.example.id
  security_groups = [aws_security_group.example.id]

  tags = {
    Name = "example-instance"
  }
}

# Create a Traffic Mirror Filter
resource "awscc_ec2_traffic_mirror_filter" "example" {
  description = "Example Traffic Mirror Filter"

  tags = [{
    key   = "Name"
    value = "example-filter"
    }, {
    key   = "Environment"
    value = "Test"
  }]
}

# Create filter rules - allow all TCP traffic
resource "awscc_ec2_traffic_mirror_filter_rule" "example_inbound" {
  traffic_mirror_filter_id = awscc_ec2_traffic_mirror_filter.example.id
  rule_number              = 100
  rule_action              = "accept"
  traffic_direction        = "ingress"
  protocol                 = 6 # TCP
  source_cidr_block        = "0.0.0.0/0"
  destination_cidr_block   = "0.0.0.0/0"
  description              = "Accept all inbound TCP traffic"
}

resource "awscc_ec2_traffic_mirror_filter_rule" "example_outbound" {
  traffic_mirror_filter_id = awscc_ec2_traffic_mirror_filter.example.id
  rule_number              = 100
  rule_action              = "accept"
  traffic_direction        = "egress"
  protocol                 = 6 # TCP
  source_cidr_block        = "0.0.0.0/0"
  destination_cidr_block   = "0.0.0.0/0"
  description              = "Accept all outbound TCP traffic"
}

# Create a Traffic Mirror Target using the Network Load Balancer
resource "awscc_ec2_traffic_mirror_target" "example" {
  network_load_balancer_arn = aws_lb.example.arn
  description               = "Example Traffic Mirror Target using NLB"

  tags = [{
    key   = "Name"
    value = "example-target"
    }, {
    key   = "Environment"
    value = "Test"
  }]

  depends_on = [aws_lb.example]
}

# Create a Traffic Mirror Session
resource "awscc_ec2_traffic_mirror_session" "example" {
  network_interface_id     = aws_instance.example.primary_network_interface_id
  traffic_mirror_filter_id = awscc_ec2_traffic_mirror_filter.example.id
  traffic_mirror_target_id = awscc_ec2_traffic_mirror_target.example.id
  session_number           = 1
  packet_length            = 100 # Mirror the first 100 bytes of each packet
  description              = "Example Traffic Mirror Session"
  virtual_network_id       = 1234

  tags = [{
    key   = "Name"
    value = "example-session"
    }, {
    key   = "Environment"
    value = "Test"
  }]

  depends_on = [
    awscc_ec2_traffic_mirror_filter.example,
    awscc_ec2_traffic_mirror_target.example,
    aws_instance.example
  ]
}
