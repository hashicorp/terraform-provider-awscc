data "aws_caller_identity" "current" {}

# Create ACM certificate for the endpoint
resource "aws_acm_certificate" "cert" {
  domain_name       = "example.com"
  validation_method = "DNS"
}

# Create a VPC for the endpoint
resource "awscc_ec2_vpc" "main" {
  cidr_block = "10.0.0.0/16"

  tags = [{
    key   = "Name"
    value = "verified-access-vpc"
  }]
}

# Create a subnet
resource "awscc_ec2_subnet" "main" {
  vpc_id     = awscc_ec2_vpc.main.id
  cidr_block = "10.0.1.0/24"

  tags = [{
    key   = "Name"
    value = "verified-access-subnet"
  }]
}

# Create security group
resource "awscc_ec2_security_group" "endpoint" {
  group_name        = "verified-access-endpoint-sg"
  group_description = "Security group for Verified Access Endpoint"
  vpc_id            = awscc_ec2_vpc.main.id

  security_group_ingress = [{
    ip_protocol = "tcp"
    from_port   = 443
    to_port     = 443
    cidr_ip     = "0.0.0.0/0"
  }]

  tags = [{
    key   = "Name"
    value = "verified-access-endpoint-sg"
  }]
}

# Create Network Load Balancer
resource "awscc_elasticloadbalancingv2_load_balancer" "nlb" {
  name    = "verified-access-nlb"
  scheme  = "internet-facing"
  type    = "network"
  subnets = [awscc_ec2_subnet.main.id]

  tags = [{
    key   = "Name"
    value = "verified-access-nlb"
  }]
}

# Create Verified Access Group
resource "awscc_ec2_verified_access_group" "example" {
  verified_access_instance_id = awscc_ec2_verified_access_instance.example.id
  description                 = "Example Verified Access Group"

  tags = [{
    key   = "Modified By"
    value = "AWSCC"
  }]
}

# Create Verified Access Instance
resource "awscc_ec2_verified_access_instance" "example" {
  description = "Example Verified Access Instance"

  tags = [{
    key   = "Modified By"
    value = "AWSCC"
  }]
}

# Policy Document
data "aws_iam_policy_document" "endpoint_policy" {
  statement {
    effect = "Allow"
    principals {
      type        = "AWS"
      identifiers = [data.aws_caller_identity.current.account_id]
    }
    actions   = ["verifiedaccess:*"]
    resources = ["*"]
  }
}

# Create Verified Access Endpoint
resource "awscc_ec2_verified_access_endpoint" "example" {
  application_domain       = "app.example.com"
  endpoint_domain_prefix   = "myapp"
  endpoint_type            = "network-interface"
  domain_certificate_arn   = aws_acm_certificate.cert.arn
  attachment_type          = "vpc"
  verified_access_group_id = awscc_ec2_verified_access_group.example.id
  description              = "Example Verified Access Endpoint"

  security_group_ids = [awscc_ec2_security_group.endpoint.id]

  network_interface_options = {
    port     = 443
    protocol = "https"
  }

  policy_enabled  = true
  policy_document = jsonencode(data.aws_iam_policy_document.endpoint_policy.json)

  tags = [{
    key   = "Modified By"
    value = "AWSCC"
  }]
}