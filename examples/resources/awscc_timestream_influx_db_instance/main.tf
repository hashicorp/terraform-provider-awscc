# Get current AWS account ID
data "aws_caller_identity" "current" {}

# Get current AWS region
data "aws_region" "current" {}

# Create VPC
resource "awscc_ec2_vpc" "example" {
  cidr_block           = "10.0.0.0/16"
  enable_dns_hostnames = true
  enable_dns_support   = true
  tags = [{
    key   = "Name"
    value = "timestream-influxdb-vpc"
  }]
}

# Create Internet Gateway
resource "awscc_ec2_internet_gateway" "example" {
  tags = [{
    key   = "Name"
    value = "timestream-influxdb-igw"
  }]
}

# Attach Internet Gateway to VPC
resource "awscc_ec2_vpc_gateway_attachment" "example" {
  vpc_id              = awscc_ec2_vpc.example.id
  internet_gateway_id = awscc_ec2_internet_gateway.example.id
}

# Create Public Subnet 1
resource "awscc_ec2_subnet" "public1" {
  vpc_id                  = awscc_ec2_vpc.example.id
  cidr_block              = "10.0.1.0/24"
  availability_zone       = "${data.aws_region.current.name}a"
  map_public_ip_on_launch = true
  tags = [{
    key   = "Name"
    value = "timestream-influxdb-subnet-1"
  }]
}

# Create Public Subnet 2
resource "awscc_ec2_subnet" "public2" {
  vpc_id                  = awscc_ec2_vpc.example.id
  cidr_block              = "10.0.2.0/24"
  availability_zone       = "${data.aws_region.current.name}b"
  map_public_ip_on_launch = true
  tags = [{
    key   = "Name"
    value = "timestream-influxdb-subnet-2"
  }]
}

# Create Route Table
resource "awscc_ec2_route_table" "public" {
  vpc_id = awscc_ec2_vpc.example.id
  tags = [{
    key   = "Name"
    value = "timestream-influxdb-rt"
  }]
}

# Create Route to Internet Gateway
resource "awscc_ec2_route" "public_internet_gateway" {
  route_table_id         = awscc_ec2_route_table.public.id
  destination_cidr_block = "0.0.0.0/0"
  gateway_id             = awscc_ec2_internet_gateway.example.id
}

# Associate Route Table with Subnet 1
resource "awscc_ec2_subnet_route_table_association" "public1" {
  subnet_id      = awscc_ec2_subnet.public1.id
  route_table_id = awscc_ec2_route_table.public.id
}

# Associate Route Table with Subnet 2
resource "awscc_ec2_subnet_route_table_association" "public2" {
  subnet_id      = awscc_ec2_subnet.public2.id
  route_table_id = awscc_ec2_route_table.public.id
}

# Create Security Group
resource "awscc_ec2_security_group" "example" {
  group_description = "Security group for InfluxDB instance"
  vpc_id            = awscc_ec2_vpc.example.id
  security_group_ingress = [
    {
      ip_protocol = "tcp"
      from_port   = 8086
      to_port     = 8086
      cidr_ip     = "0.0.0.0/0"
    }
  ]
  tags = [{
    key   = "Name"
    value = "timestream-influxdb-sg"
  }]
}

# Create the InfluxDB instance
resource "awscc_timestream_influx_db_instance" "example" {
  name              = "example-influxdb"
  db_instance_type  = "db.influx.medium"
  allocated_storage = 200
  db_storage_type   = "InfluxIOIncludedT1"
  deployment_type   = "WITH_MULTIAZ_STANDBY"
  bucket            = "example-bucket"
  organization      = "example-org"
  username          = "admin"
  password          = "Example123" # Change this in production

  publicly_accessible    = true
  vpc_subnet_ids         = [awscc_ec2_subnet.public1.id, awscc_ec2_subnet.public2.id]
  vpc_security_group_ids = [awscc_ec2_security_group.example.id]

  log_delivery_configuration = {
    s3_configuration = {
      bucket_name = "example-logs-${data.aws_caller_identity.current.account_id}"
      enabled     = false
    }
  }

  tags = [{
    key   = "Environment"
    value = "Example"
  }]
}