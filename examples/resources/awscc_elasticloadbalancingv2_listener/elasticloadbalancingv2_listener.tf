# Create VPC for the load balancer
resource "awscc_ec2_vpc" "main" {
  cidr_block = "10.0.0.0/16"
  tags = [
    {
      key   = "Name"
      value = "vpc-for-lb"
    }
  ]
}

# Internet Gateway
resource "awscc_ec2_internet_gateway" "main" {
  tags = [
    {
      key   = "Name"
      value = "igw-for-lb"
    }
  ]
}

# Attach Internet Gateway to VPC
resource "awscc_ec2_vpc_gateway_attachment" "main" {
  internet_gateway_id = awscc_ec2_internet_gateway.main.id
  vpc_id              = awscc_ec2_vpc.main.id
}

# Route Table
resource "awscc_ec2_route_table" "public" {
  vpc_id = awscc_ec2_vpc.main.id
  tags = [
    {
      key   = "Name"
      value = "public-rt"
    }
  ]
}

# Route to Internet Gateway
resource "awscc_ec2_route" "public" {
  destination_cidr_block = "0.0.0.0/0"
  route_table_id         = awscc_ec2_route_table.public.id
  gateway_id             = awscc_ec2_internet_gateway.main.id
  depends_on             = [awscc_ec2_vpc_gateway_attachment.main]
}

# Create subnets for the load balancer
resource "awscc_ec2_subnet" "subnet1" {
  vpc_id            = awscc_ec2_vpc.main.id
  cidr_block        = "10.0.1.0/24"
  availability_zone = "us-west-2a"
  tags = [
    {
      key   = "Name"
      value = "lb-subnet-1"
    }
  ]
}

resource "awscc_ec2_subnet" "subnet2" {
  vpc_id            = awscc_ec2_vpc.main.id
  cidr_block        = "10.0.2.0/24"
  availability_zone = "us-west-2b"
  tags = [
    {
      key   = "Name"
      value = "lb-subnet-2"
    }
  ]
}

# Route Table Associations
resource "awscc_ec2_subnet_route_table_association" "subnet1" {
  route_table_id = awscc_ec2_route_table.public.id
  subnet_id      = awscc_ec2_subnet.subnet1.id
}

resource "awscc_ec2_subnet_route_table_association" "subnet2" {
  route_table_id = awscc_ec2_route_table.public.id
  subnet_id      = awscc_ec2_subnet.subnet2.id
}

# Security group for the load balancer
resource "awscc_ec2_security_group" "lb_sg" {
  vpc_id            = awscc_ec2_vpc.main.id
  group_description = "Security group for load balancer"
  security_group_ingress = [
    {
      ip_protocol = "tcp"
      from_port   = 80
      to_port     = 80
      cidr_ip     = "0.0.0.0/0"
    }
  ]
  tags = [
    {
      key   = "Name"
      value = "lb-security-group"
    }
  ]
}

# Load balancer
resource "awscc_elasticloadbalancingv2_load_balancer" "example" {
  name            = "example-lb"
  security_groups = [awscc_ec2_security_group.lb_sg.id]
  subnets = [
    awscc_ec2_subnet.subnet1.id,
    awscc_ec2_subnet.subnet2.id
  ]
  scheme          = "internet-facing"
  type            = "application"
  ip_address_type = "ipv4"
  tags = [
    {
      key   = "Name"
      value = "example-lb"
    }
  ]
  depends_on = [awscc_ec2_vpc_gateway_attachment.main]
}

# Target group
resource "awscc_elasticloadbalancingv2_target_group" "example" {
  name                  = "example-target-group"
  port                  = 80
  protocol              = "HTTP"
  vpc_id                = awscc_ec2_vpc.main.id
  target_type           = "ip"
  health_check_enabled  = true
  health_check_path     = "/"
  health_check_protocol = "HTTP"
  matcher = {
    http_code = "200"
  }
  tags = [
    {
      key   = "Name"
      value = "example-target-group"
    }
  ]
}

# Example HTTP Listener with forward action
resource "awscc_elasticloadbalancingv2_listener" "example" {
  load_balancer_arn = awscc_elasticloadbalancingv2_load_balancer.example.load_balancer_arn
  port              = 80
  protocol          = "HTTP"

  default_actions = [
    {
      type             = "forward"
      target_group_arn = awscc_elasticloadbalancingv2_target_group.example.target_group_arn
    }
  ]
}
