# Get current AWS region
data "aws_region" "current" {}

# Get current AWS account ID
data "aws_caller_identity" "current" {}

# Create a VPC for QuickSight VPC connection
resource "awscc_ec2_vpc" "quicksight_vpc" {
  cidr_block = "10.0.0.0/16"
  tags = [{
    key   = "Name"
    value = "quicksight-vpc"
    }, {
    key   = "Modified By"
    value = "AWSCC"
  }]
}

# Create subnets
resource "awscc_ec2_subnet" "quicksight_subnet_1" {
  vpc_id            = awscc_ec2_vpc.quicksight_vpc.id
  cidr_block        = "10.0.1.0/24"
  availability_zone = "${data.aws_region.current.name}a"
  tags = [{
    key   = "Name"
    value = "quicksight-subnet-1"
    }, {
    key   = "Modified By"
    value = "AWSCC"
  }]
}

resource "awscc_ec2_subnet" "quicksight_subnet_2" {
  vpc_id            = awscc_ec2_vpc.quicksight_vpc.id
  cidr_block        = "10.0.2.0/24"
  availability_zone = "${data.aws_region.current.name}b"
  tags = [{
    key   = "Name"
    value = "quicksight-subnet-2"
    }, {
    key   = "Modified By"
    value = "AWSCC"
  }]
}

# Create security group
resource "awscc_ec2_security_group" "quicksight_sg" {
  group_description = "Security group for QuickSight VPC connection"
  vpc_id            = awscc_ec2_vpc.quicksight_vpc.id
  group_name        = "quicksight-vpc-sg"
  security_group_ingress = [{
    ip_protocol = "-1"
    from_port   = 0
    to_port     = 0
    cidr_ipv4   = "10.0.0.0/16"
  }]
  tags = [{
    key   = "Name"
    value = "quicksight-sg"
    }, {
    key   = "Modified By"
    value = "AWSCC"
  }]
}

# Create IAM role for QuickSight VPC connection
data "aws_iam_policy_document" "quicksight_vpc_assume_role" {
  statement {
    actions = ["sts:AssumeRole"]
    effect  = "Allow"
    principals {
      type        = "Service"
      identifiers = ["quicksight.amazonaws.com"]
    }
  }
}

data "aws_iam_policy_document" "quicksight_vpc_policy" {
  statement {
    effect = "Allow"
    actions = [
      "ec2:CreateNetworkInterface",
      "ec2:DeleteNetworkInterface",
      "ec2:DescribeNetworkInterfaces",
      "ec2:DescribeSubnets",
      "ec2:DescribeSecurityGroups",
      "ec2:DescribeVpcs"
    ]
    resources = ["*"]
  }
}

resource "aws_iam_role" "quicksight_vpc_role" {
  name               = "quicksight-vpc-role"
  assume_role_policy = data.aws_iam_policy_document.quicksight_vpc_assume_role.json

  inline_policy {
    name   = "quicksight-vpc-policy"
    policy = data.aws_iam_policy_document.quicksight_vpc_policy.json
  }

  tags = {
    Modified_By = "AWSCC"
  }
}

# Create QuickSight VPC connection
resource "awscc_quicksight_vpc_connection" "example" {
  aws_account_id     = data.aws_caller_identity.current.account_id
  vpc_connection_id  = "example-vpc-connection"
  name               = "Example VPC Connection"
  role_arn           = aws_iam_role.quicksight_vpc_role.arn
  security_group_ids = [awscc_ec2_security_group.quicksight_sg.id]
  subnet_ids         = [awscc_ec2_subnet.quicksight_subnet_1.id, awscc_ec2_subnet.quicksight_subnet_2.id]
  tags = [{
    key   = "Modified By"
    value = "AWSCC"
  }]
}