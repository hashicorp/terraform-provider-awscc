# Create a VPC first
resource "awscc_ec2_vpc" "example" {
  cidr_block = "10.0.0.0/16"
  tags = [{
    key   = "Modified By"
    value = "AWSCC"
  }]
}

# Add secondary CIDR block to the VPC
resource "awscc_ec2_vpc_cidr_block" "secondary_cidr" {
  vpc_id     = awscc_ec2_vpc.example.id
  cidr_block = "10.1.0.0/16"
}