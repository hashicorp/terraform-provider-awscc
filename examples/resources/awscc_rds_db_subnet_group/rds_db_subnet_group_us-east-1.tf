resource "awscc_ec2_vpc" "main" {
  cidr_block = "10.0.0.0/16"
}

resource "awscc_ec2_subnet" "subnet_a" {
  vpc_id            = resource.awscc_ec2_vpc.main.id
  cidr_block        = "10.0.1.0/24"
  availability_zone = "us-east-1a"
}

resource "awscc_ec2_subnet" "subnet_b" {
  vpc_id            = resource.awscc_ec2_vpc.main.id
  cidr_block        = "10.0.2.0/24"
  availability_zone = "us-east-1b"
}

resource "awscc_rds_db_subnet_group" "example" {
  db_subnet_group_description = "example description"
  db_subnet_group_name        = "example-subnet-group"
  subnet_ids                  = [awscc_ec2_subnet.subnet_a.subnet_id, awscc_ec2_subnet.subnet_b.subnet_id]
}