
# Create primary VPC
resource "awscc_ec2_vpc" "primary" {
  cidr_block           = "10.0.0.0/16"
  enable_dns_hostnames = true
  enable_dns_support   = true

  tags = [{
    key   = "Modified By"
    value = "AWSCC"
  }]
}

# Create secondary VPC
resource "awscc_ec2_vpc" "secondary" {
  cidr_block           = "172.16.0.0/16"
  enable_dns_hostnames = true
  enable_dns_support   = true

  tags = [{
    key   = "Modified By"
    value = "AWSCC"
  }]
}

# Create a security group in the primary VPC
resource "awscc_ec2_security_group" "example" {
  group_description = "Example security group"
  group_name        = "example-security-group"
  vpc_id            = awscc_ec2_vpc.primary.id

  security_group_ingress = [{
    ip_protocol = "tcp"
    from_port   = 80
    to_port     = 80
    cidr_ip     = "0.0.0.0/0"
  }]

  tags = [{
    key   = "Modified By"
    value = "AWSCC"
  }]
}

# Associate the security group with the secondary VPC
resource "awscc_ec2_security_group_vpc_association" "example" {
  group_id = awscc_ec2_security_group.example.id
  vpc_id   = awscc_ec2_vpc.secondary.id
}