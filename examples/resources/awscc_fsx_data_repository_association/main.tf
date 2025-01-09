data "aws_region" "current" {}

data "aws_caller_identity" "current" {}

# VPC for FSx deployment
resource "awscc_ec2_vpc" "example" {
  cidr_block           = "10.0.0.0/16"
  enable_dns_hostnames = true
  enable_dns_support   = true

  tags = [{
    key   = "Modified By"
    value = "AWSCC"
  }]
}

resource "awscc_ec2_subnet" "example" {
  vpc_id            = awscc_ec2_vpc.example.id
  cidr_block        = "10.0.1.0/24"
  availability_zone = "${data.aws_region.current.name}a"

  tags = [{
    key   = "Modified By"
    value = "AWSCC"
  }]
}

resource "awscc_ec2_security_group" "example" {
  group_description = "FSx security group"
  vpc_id            = awscc_ec2_vpc.example.id

  tags = [{
    key   = "Modified By"
    value = "AWSCC"
  }]
}

# S3 bucket for FSx data repository
resource "awscc_s3_bucket" "fsx_data" {
  bucket_name = "fsx-dra-example-${data.aws_caller_identity.current.account_id}-${data.aws_region.current.name}"

  tags = [{
    key   = "Modified By"
    value = "AWSCC"
  }]
}

# FSx Lustre File System
resource "aws_fsx_lustre_file_system" "example" {
  storage_capacity            = 1200
  subnet_ids                  = [awscc_ec2_subnet.example.id]
  security_group_ids          = [awscc_ec2_security_group.example.id]
  deployment_type             = "PERSISTENT_1"
  per_unit_storage_throughput = 50

  tags = {
    "Modified By" = "AWSCC"
  }
}

# FSx Data Repository Association
resource "awscc_fsx_data_repository_association" "example" {
  file_system_id       = aws_fsx_lustre_file_system.example.id
  file_system_path     = "/data"
  data_repository_path = "s3://${awscc_s3_bucket.fsx_data.bucket_name}/data"

  batch_import_meta_data_on_create = true
  imported_file_chunk_size         = 1024

  s3 = {
    auto_import_policy = {
      events = ["NEW", "CHANGED", "DELETED"]
    }
    auto_export_policy = {
      events = ["NEW", "CHANGED", "DELETED"]
    }
  }

  tags = [{
    key   = "Modified By"
    value = "AWSCC"
  }]
}