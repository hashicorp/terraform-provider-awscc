# Get current region
data "aws_region" "current" {}

# Create VPC for SimpleAD
resource "aws_vpc" "simple_ad_vpc" {
  cidr_block = "10.0.0.0/16"
  tags = {
    "Modified By" = "AWSCC"
  }
}

# Create internet gateway
resource "aws_internet_gateway" "simple_ad_igw" {
  vpc_id = aws_vpc.simple_ad_vpc.id
  tags = {
    "Modified By" = "AWSCC"
  }
}

# Create subnets in different AZs
resource "aws_subnet" "simple_ad_subnet_1" {
  vpc_id            = aws_vpc.simple_ad_vpc.id
  cidr_block        = "10.0.1.0/24"
  availability_zone = "${data.aws_region.current.name}a"
  tags = {
    "Modified By" = "AWSCC"
  }
}

resource "aws_subnet" "simple_ad_subnet_2" {
  vpc_id            = aws_vpc.simple_ad_vpc.id
  cidr_block        = "10.0.2.0/24"
  availability_zone = "${data.aws_region.current.name}b"
  tags = {
    "Modified By" = "AWSCC"
  }
}

# Create Simple AD
resource "awscc_directoryservice_simple_ad" "example" {
  name        = "corp.example.com"
  password    = "SuperSecretPassw0rd"
  size        = "Small"
  description = "Simple AD example directory"
  vpc_settings = {
    vpc_id     = aws_vpc.simple_ad_vpc.id
    subnet_ids = [aws_subnet.simple_ad_subnet_1.id, aws_subnet.simple_ad_subnet_2.id]
  }
}