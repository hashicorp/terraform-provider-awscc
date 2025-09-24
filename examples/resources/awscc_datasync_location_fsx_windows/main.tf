# Get current AWS region and account ID
data "aws_region" "current" {}
data "aws_caller_identity" "current" {}

# Create a VPC for FSx Windows File System
resource "awscc_ec2_vpc" "fsx" {
  cidr_block = "10.0.0.0/16"
  tags = [{
    key   = "Name"
    value = "FSx-VPC"
  }]
}

# Create subnets
resource "awscc_ec2_subnet" "fsx" {
  vpc_id            = awscc_ec2_vpc.fsx.id
  cidr_block        = "10.0.1.0/24"
  availability_zone = "${data.aws_region.current.name}a"
  tags = [{
    key   = "Name"
    value = "FSx-Subnet"
  }]
}

# Create security group for FSx
resource "awscc_ec2_security_group" "fsx" {
  group_description = "Security group for FSx Windows File System"
  vpc_id            = awscc_ec2_vpc.fsx.id
  tags = [{
    key   = "Name"
    value = "fsx-sg"
  }]
}

# Create FSx Windows File System
resource "aws_fsx_windows_file_system" "example" {
  storage_capacity    = 32
  subnet_ids          = [awscc_ec2_subnet.fsx.id]
  throughput_capacity = 16
  security_group_ids  = [awscc_ec2_security_group.fsx.id]

  self_managed_active_directory {
    dns_ips     = ["10.0.1.10"]
    domain_name = "example.com"
    password    = "Password123!"
    username    = "Admin"
  }
}

# Create DataSync location for FSx Windows
resource "awscc_datasync_location_fsx_windows" "example" {
  fsx_filesystem_arn  = aws_fsx_windows_file_system.example.arn
  security_group_arns = ["arn:aws:ec2:${data.aws_region.current.name}:${data.aws_caller_identity.current.account_id}:security-group/${awscc_ec2_security_group.fsx.id}"]
  user                = "Admin"
  password            = "Password123!"
  domain              = "example.com"
  subdirectory        = "/share"

  tags = [{
    key   = "Modified By"
    value = "AWSCC"
  }]
}

# Output the DataSync location ARN
output "datasync_location_arn" {
  value = awscc_datasync_location_fsx_windows.example.location_arn
}