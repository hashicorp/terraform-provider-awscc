data "aws_region" "current" {}

# VPC and Networking Resources
resource "awscc_ec2_vpc" "example" {
  cidr_block = "10.0.0.0/16"
  tags = [{
    key   = "Name"
    value = "mediaconnect-example-vpc"
  }]
}

resource "awscc_ec2_subnet" "example" {
  vpc_id            = awscc_ec2_vpc.example.id
  cidr_block        = "10.0.1.0/24"
  availability_zone = "${data.aws_region.current.name}a"
  tags = [{
    key   = "Name"
    value = "mediaconnect-example-subnet"
  }]
}

resource "awscc_ec2_security_group" "example" {
  vpc_id            = awscc_ec2_vpc.example.id
  group_description = "Security group for MediaConnect VPC Interface"
  tags = [{
    key   = "Name"
    value = "mediaconnect-example-sg"
  }]
}

# IAM Role for MediaConnect
data "aws_iam_policy_document" "assume_role" {
  statement {
    effect = "Allow"
    principals {
      type        = "Service"
      identifiers = ["mediaconnect.amazonaws.com"]
    }
    actions = ["sts:AssumeRole"]
  }
}

data "aws_iam_policy_document" "vpc_interface" {
  statement {
    effect = "Allow"
    actions = [
      "ec2:CreateNetworkInterface",
      "ec2:DescribeNetworkInterfaces",
      "ec2:DescribeSubnets",
      "ec2:DeleteNetworkInterface"
    ]
    resources = ["*"]
  }
}

resource "awscc_iam_role" "mediaconnect" {
  assume_role_policy_document = jsonencode(jsondecode(data.aws_iam_policy_document.assume_role.json))
  description                 = "Role for MediaConnect VPC Interface"
  role_name                   = "MediaConnectVPCInterface"
  managed_policy_arns         = []
  tags = [{
    key   = "Modified By"
    value = "AWSCC"
  }]
}

resource "awscc_iam_role_policy" "mediaconnect" {
  policy_document = jsonencode(jsondecode(data.aws_iam_policy_document.vpc_interface.json))
  policy_name     = "MediaConnectVPCInterfacePolicy"
  role_name       = awscc_iam_role.mediaconnect.role_name
}

# MediaConnect Flow
resource "awscc_mediaconnect_flow" "example" {
  name              = "example-flow"
  availability_zone = "${data.aws_region.current.name}a"
  source = {
    name               = "example-source"
    description        = "Example source"
    protocol           = "zixi-push"
    ingest_port        = 2088
    max_bit_rate       = 10000000
    min_latency        = 2000
    vpc_interface_name = "example-vpc-interface"
    whitelist_cidr     = "0.0.0.0/0"
  }
}

# MediaConnect Flow VPC Interface
resource "awscc_mediaconnect_flow_vpc_interface" "example" {
  flow_arn           = awscc_mediaconnect_flow.example.flow_arn
  name               = "example-vpc-interface"
  role_arn           = awscc_iam_role.mediaconnect.arn
  security_group_ids = [awscc_ec2_security_group.example.id]
  subnet_id          = awscc_ec2_subnet.example.id
}