#Basic

resource "awscc_ec2_vpc" "main" {
  cidr_block = "10.0.0.0/16"
}

resource "awscc_ec2_vpc_endpoint" "s3" {
  vpc_id       = awscc_ec2_vpc.main.id
  service_name = "com.amazonaws.us-west-2.s3"
}