resource "awscc_efs_mount_target" "main" {
  file_system_id  = awscc_efs_file_system.main.id
  subnet_id       = awscc_ec2_subnet.main.id
  security_groups = ["sg-xxxxxx"]
}

resource "awscc_efs_file_system" "main" {

  file_system_tags = [
    {
      key   = "Name"
      value = "this"
    },
    {
      key   = "Modified By"
      value = "AWSCC"
    }
  ]
}

resource "awscc_ec2_vpc" "main" {
  cidr_block = "10.0.0.0/16"
}

resource "awscc_ec2_subnet" "main" {
  vpc_id     = resource.awscc_ec2_vpc.main.id
  cidr_block = "10.0.1.0/24"
  tags = [{
    key   = "Modified By"
    value = "AWSCC"
  }]
}