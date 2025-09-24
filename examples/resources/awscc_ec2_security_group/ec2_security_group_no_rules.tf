resource "awscc_ec2_security_group" "example" {
  group_description = "Security group example"
  vpc_id            = awscc_ec2_vpc.selected.id

  tags = [
    {
      key   = "Name"
      value = "Example SG"
    }
  ]
}

resource "awscc_ec2_vpc" "selected" {
  cidr_block = "10.0.0.0/16"
}