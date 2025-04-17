# Get default VPC
data "aws_vpc" "default" {
  default = true
}

data "aws_subnets" "default" {
  filter {
    name   = "vpc-id"
    values = [data.aws_vpc.default.id]
  }
}

data "aws_subnet" "selected" {
  id = tolist(data.aws_subnets.default.ids)[0]
}

# Create Security Groups
resource "awscc_ec2_security_group" "example_1" {
  group_description = "Security group for source instance"
  vpc_id            = data.aws_vpc.default.id
  security_group_ingress = [{
    description = "Allow SSH"
    from_port   = 22
    to_port     = 22
    ip_protocol = "tcp"
    cidr_ip     = "0.0.0.0/0"
  }]
  tags = [{
    key   = "Name"
    value = "network-insights-sg-1"
  }]
}

resource "awscc_ec2_security_group" "example_2" {
  group_description = "Security group for destination instance"
  vpc_id            = data.aws_vpc.default.id
  security_group_ingress = [{
    description = "Allow HTTP"
    from_port   = 80
    to_port     = 80
    ip_protocol = "tcp"
    cidr_ip     = "0.0.0.0/0"
  }]
  tags = [{
    key   = "Name"
    value = "network-insights-sg-2"
  }]
}

data "aws_ami" "amazon_linux_2" {
  most_recent = true

  filter {
    name   = "name"
    values = ["amzn2-ami-hvm-*-x86_64-gp2"]
  }

  filter {
    name   = "virtualization-type"
    values = ["hvm"]
  }

  owners = ["amazon"]
}

# Create EC2 Instances
resource "awscc_ec2_instance" "source" {
  instance_type = "t3.micro"
  subnet_id     = data.aws_subnet.selected.id
  security_group_ids = [
    awscc_ec2_security_group.example_1.id
  ]
  image_id = data.aws_ami.amazon_linux_2.id
  tags = [{
    key   = "Name"
    value = "network-insights-source"
  }]
}

resource "awscc_ec2_instance" "destination" {
  instance_type = "t3.micro"
  subnet_id     = data.aws_subnet.selected.id
  security_group_ids = [
    awscc_ec2_security_group.example_2.id
  ]
  image_id = data.aws_ami.amazon_linux_2.id
  tags = [{
    key   = "Name"
    value = "network-insights-destination"
  }]
}

# Create Network Insights Path
resource "awscc_ec2_network_insights_path" "example" {
  source      = awscc_ec2_instance.source.id
  destination = awscc_ec2_instance.destination.id
  protocol    = "tcp"

  destination_port = 80

  tags = [{
    key   = "Name"
    value = "example-network-insights-path"
    }, {
    key   = "Modified By"
    value = "AWSCC"
  }]
}