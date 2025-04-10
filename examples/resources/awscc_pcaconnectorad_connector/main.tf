# Get current region
data "aws_region" "current" {}

# Create VPC for Directory Service
resource "awscc_ec2_vpc" "main" {
  cidr_block           = "10.0.0.0/16"
  enable_dns_support   = true
  enable_dns_hostnames = true
  tags = [{
    key   = "Name"
    value = "pca-connector-vpc"
  }]
}

# Create subnets for Directory Service
resource "awscc_ec2_subnet" "subnet1" {
  vpc_id            = awscc_ec2_vpc.main.id
  cidr_block        = "10.0.1.0/24"
  availability_zone = "${data.aws_region.current.name}a"
  tags = [{
    key   = "Name"
    value = "pca-connector-subnet1"
  }]
}

resource "awscc_ec2_subnet" "subnet2" {
  vpc_id            = awscc_ec2_vpc.main.id
  cidr_block        = "10.0.2.0/24"
  availability_zone = "${data.aws_region.current.name}b"
  tags = [{
    key   = "Name"
    value = "pca-connector-subnet2"
  }]
}

# Create security group for Directory Service
resource "awscc_ec2_security_group" "directory" {
  group_description = "Security group for AWS Managed Microsoft AD"
  vpc_id            = awscc_ec2_vpc.main.id
  tags = [{
    key   = "Name"
    value = "pca-connector-directory-sg"
  }]
}

# Allow inbound traffic for Directory Service
resource "awscc_ec2_security_group_ingress" "directory" {
  group_id    = awscc_ec2_security_group.directory.id
  ip_protocol = "tcp"
  from_port   = 389
  to_port     = 389
  cidr_ip     = "10.0.0.0/16"
  description = "Allow LDAP"
}

# Create AWS Managed Microsoft AD using AWS provider
resource "aws_directory_service_directory" "ad" {
  name     = "pca.example.com"
  password = "SuperSecretPassw0rd1!"
  edition  = "Standard"
  type     = "MicrosoftAD"

  vpc_settings {
    vpc_id = awscc_ec2_vpc.main.id
    subnet_ids = [
      awscc_ec2_subnet.subnet1.id,
      awscc_ec2_subnet.subnet2.id
    ]
  }
}

# Create Private Certificate Authority
resource "awscc_acmpca_certificate_authority" "ca" {
  key_algorithm     = "RSA_2048"
  signing_algorithm = "SHA256WITHRSA"
  subject = {
    common_name         = "pca.example.com"
    country             = "US"
    organization        = "Example Corp"
    organizational_unit = "IT"
    state               = "WA"
    locality            = "Seattle"
  }
  type = "ROOT"
  tags = [{
    key   = "Name"
    value = "Example PCA"
  }]
}

# Create PCA Connector AD
resource "awscc_pcaconnectorad_connector" "example" {
  certificate_authority_arn = awscc_acmpca_certificate_authority.ca.arn
  directory_id              = aws_directory_service_directory.ad.id
  vpc_information = {
    security_group_ids = [awscc_ec2_security_group.directory.id]
  }
  tags = [{
    key   = "Environment"
    value = "test"
    },
    {
      key   = "Project"
      value = "pca-connector-example"
  }]
}