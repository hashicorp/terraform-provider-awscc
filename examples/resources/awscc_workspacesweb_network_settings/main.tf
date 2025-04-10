# Create VPC
resource "awscc_ec2_vpc" "example" {
  cidr_block = "10.0.0.0/16"
  tags = [{
    key   = "Name"
    value = "workspaces-web-example"
  }]
}

# Create Internet Gateway
resource "awscc_ec2_internet_gateway" "example" {
  tags = [{
    key   = "Name"
    value = "workspaces-web-example"
  }]
}

# Attach Internet Gateway to VPC
resource "awscc_ec2_vpc_gateway_attachment" "example" {
  vpc_id              = awscc_ec2_vpc.example.id
  internet_gateway_id = awscc_ec2_internet_gateway.example.id
}

# Create Subnets
resource "awscc_ec2_subnet" "example_1" {
  vpc_id     = awscc_ec2_vpc.example.id
  cidr_block = "10.0.1.0/24"
  tags = [{
    key   = "Name"
    value = "workspaces-web-example-1"
  }]
}

resource "awscc_ec2_subnet" "example_2" {
  vpc_id     = awscc_ec2_vpc.example.id
  cidr_block = "10.0.2.0/24"
  tags = [{
    key   = "Name"
    value = "workspaces-web-example-2"
  }]
}

# Create Security Group
resource "awscc_ec2_security_group" "example" {
  group_description = "Security group for Workspaces Web Network Settings"
  vpc_id            = awscc_ec2_vpc.example.id
  tags = [{
    key   = "Name"
    value = "workspaces-web-example"
  }]
}

# Workspaces Web Network Settings
resource "awscc_workspacesweb_network_settings" "example" {
  vpc_id             = awscc_ec2_vpc.example.id
  subnet_ids         = [awscc_ec2_subnet.example_1.id, awscc_ec2_subnet.example_2.id]
  security_group_ids = [awscc_ec2_security_group.example.id]

  tags = [{
    key   = "Modified By"
    value = "AWSCC"
  }]
}