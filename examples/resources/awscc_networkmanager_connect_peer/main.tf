data "aws_region" "current" {}
data "aws_caller_identity" "current" {}

# Create VPC and Subnet for the Connect Peer
resource "awscc_ec2_vpc" "example" {
  cidr_block           = "10.0.0.0/16"
  enable_dns_hostnames = true
  enable_dns_support   = true

  tags = [{
    key   = "Name"
    value = "connect-peer-example-vpc"
  }]
}

resource "awscc_ec2_subnet" "example" {
  vpc_id                  = awscc_ec2_vpc.example.id
  cidr_block              = "10.0.1.0/24"
  availability_zone       = "${data.aws_region.current.name}a"
  map_public_ip_on_launch = false

  tags = [{
    key   = "Name"
    value = "connect-peer-example-subnet"
  }]
}

# Create the Global Network and Core Network
resource "awscc_networkmanager_global_network" "example" {
  description = "Example Global Network for Connect Peer"

  tags = [{
    key   = "Name"
    value = "connect-peer-example-global-network"
  }]
}

# Core Network policy document
data "aws_iam_policy_document" "core_network_policy" {
  statement {
    effect = "Allow"
    actions = [
      "network-manager:*",
      "ec2:DescribeVpcs",
      "ec2:DescribeSubnets"
    ]
    resources = ["*"]
  }
}

resource "awscc_networkmanager_core_network" "example" {
  global_network_id = awscc_networkmanager_global_network.example.id
  policy_document   = jsonencode(jsondecode(data.aws_iam_policy_document.core_network_policy.json))
  description       = "Example Core Network for Connect Peer"

  tags = [{
    key   = "Name"
    value = "connect-peer-example-core-network"
  }]
}

# Create VPC attachment
resource "awscc_networkmanager_vpc_attachment" "example" {
  core_network_id = awscc_networkmanager_core_network.example.core_network_id
  vpc_arn         = "arn:aws:ec2:${data.aws_region.current.name}:${data.aws_caller_identity.current.account_id}:vpc/${awscc_ec2_vpc.example.id}"
  subnet_arns     = ["arn:aws:ec2:${data.aws_region.current.name}:${data.aws_caller_identity.current.account_id}:subnet/${awscc_ec2_subnet.example.id}"]

  tags = [{
    key   = "Name"
    value = "connect-peer-example-vpc-attachment"
  }]
}

# Create Connect Attachment
resource "awscc_networkmanager_connect_attachment" "example" {
  core_network_id         = awscc_networkmanager_core_network.example.core_network_id
  transport_attachment_id = awscc_networkmanager_vpc_attachment.example.attachment_id
  edge_location           = data.aws_region.current.name
  options = {
    protocol = "GRE"
  }

  tags = [{
    key   = "Name"
    value = "connect-peer-example-connect-attachment"
  }]
}

# Create Connect Peer
resource "awscc_networkmanager_connect_peer" "example" {
  connect_attachment_id = awscc_networkmanager_connect_attachment.example.attachment_id
  peer_address          = "10.0.0.1"
  bgp_options = {
    peer_asn = 65000
  }
  inside_cidr_blocks = ["169.254.6.0/29"]
  subnet_arn         = "arn:aws:ec2:${data.aws_region.current.name}:${data.aws_caller_identity.current.account_id}:subnet/${awscc_ec2_subnet.example.id}"

  tags = [{
    key   = "Name"
    value = "connect-peer-example"
  }]
}