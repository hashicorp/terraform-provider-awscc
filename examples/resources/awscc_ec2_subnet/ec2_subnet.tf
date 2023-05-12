resource "awscc_ec2_subnet" "main" {
  vpc_id     = resource.awscc_ec2_vpc.main.id
  cidr_block = "10.0.1.0/24"
  tags = [{
    key   = "Name"
    value = "main"
  }]
}

