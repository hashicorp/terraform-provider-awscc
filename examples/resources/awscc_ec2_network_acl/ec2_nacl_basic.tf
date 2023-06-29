resource "awscc_ec2_vpc" "main" {
  cidr_block = "10.0.0.0/16"
}

resource "awscc_ec2_network_acl" "main" {
  vpc_id = awscc_ec2_vpc.main.vpc_id
  tags = [
    {
      key   = "Name"
      value = "SampleNACL"
    },
    {
      key   = "Environment"
      value = "Development"
    },
    {
      key   = "Modified By"
      value = "AWSCC"
    }
  ]
}