# Get current AWS account ID and region
data "aws_caller_identity" "current" {}
data "aws_region" "current" {}

# Create VPC and Subnet
resource "awscc_ec2_vpc" "example" {
  cidr_block = "10.0.0.0/16"
  tags = [{
    key   = "Name"
    value = "DataSync-EFS-Example"
  }]
}

resource "awscc_ec2_subnet" "example" {
  vpc_id     = awscc_ec2_vpc.example.id
  cidr_block = "10.0.1.0/24"
  tags = [{
    key   = "Name"
    value = "DataSync-EFS-Example"
  }]
}

# Create Security Group for EFS
resource "awscc_ec2_security_group" "example" {
  group_name        = "datasync-efs-sg"
  group_description = "Security group for EFS mount target"
  vpc_id            = awscc_ec2_vpc.example.id

  security_group_ingress = [
    {
      ip_protocol = "tcp"
      from_port   = 2049
      to_port     = 2049
      cidr_ip     = "10.0.0.0/16"
    }
  ]

  tags = [{
    key   = "Name"
    value = "DataSync-EFS-Example"
  }]
}

# Create EFS File System
resource "awscc_efs_file_system" "example" {
  encrypted = true
}

# Create EFS Mount Target
resource "awscc_efs_mount_target" "example" {
  file_system_id  = awscc_efs_file_system.example.id
  subnet_id       = awscc_ec2_subnet.example.id
  security_groups = [awscc_ec2_security_group.example.id]
}

# IAM role for DataSync
data "aws_iam_policy_document" "assume_role" {
  statement {
    actions = ["sts:AssumeRole"]
    principals {
      type        = "Service"
      identifiers = ["datasync.amazonaws.com"]
    }
  }
}

# Create IAM role
resource "awscc_iam_role" "datasync" {
  assume_role_policy_document = jsonencode(data.aws_iam_policy_document.assume_role.json)
  description                 = "Role for DataSync to access EFS"
  path                        = "/"
  role_name                   = "DataSyncEFSRole"
  tags = [{
    key   = "Modified By"
    value = "AWSCC"
  }]
}

# Create DataSync Location EFS
resource "awscc_datasync_location_efs" "example" {
  ec_2_config = {
    security_group_arns = ["arn:aws:ec2:${data.aws_region.current.name}:${data.aws_caller_identity.current.account_id}:security-group/${awscc_ec2_security_group.example.id}"]
    subnet_arn          = "arn:aws:ec2:${data.aws_region.current.name}:${data.aws_caller_identity.current.account_id}:subnet/${awscc_ec2_subnet.example.id}"
  }

  efs_filesystem_arn          = awscc_efs_file_system.example.arn
  file_system_access_role_arn = awscc_iam_role.datasync.arn
  subdirectory                = "/example"
  in_transit_encryption       = "TLS1_2"

  tags = [{
    key   = "Modified By"
    value = "AWSCC"
  }]
}