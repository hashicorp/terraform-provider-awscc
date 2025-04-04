# Create a VPC for testing
resource "awscc_ec2_vpc" "example" {
  cidr_block = "10.0.0.0/16"
  tags = [{
    key   = "Name"
    value = "example-vpc"
    }, {
    key   = "Modified By"
    value = "AWSCC"
  }]
}

# Create VPC Block Public Access Exclusion
resource "awscc_ec2_vpc_block_public_access_exclusion" "example" {
  vpc_id                          = awscc_ec2_vpc.example.id
  internet_gateway_exclusion_mode = "allow-bidirectional"
  tags = [{
    key   = "Name"
    value = "example-vpc-bpa-exclusion"
    }, {
    key   = "Modified By"
    value = "AWSCC"
  }]
}