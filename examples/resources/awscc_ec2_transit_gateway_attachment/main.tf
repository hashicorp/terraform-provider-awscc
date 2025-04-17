# Create a VPC for testing
resource "awscc_ec2_vpc" "example" {
  cidr_block = "10.0.0.0/16"

  tags = [{
    key   = "Name"
    value = "TGW-Attachment-Example"
  }]
}

# Create subnets for the TGW attachment
resource "awscc_ec2_subnet" "example" {
  vpc_id     = awscc_ec2_vpc.example.id
  cidr_block = "10.0.1.0/24"

  tags = [{
    key   = "Name"
    value = "TGW-Attachment-Subnet"
  }]
}

# Create a Transit Gateway
resource "awscc_ec2_transit_gateway" "example" {
  tags = [{
    key   = "Name"
    value = "Example-TGW"
  }]
}

# Create the Transit Gateway Attachment
resource "awscc_ec2_transit_gateway_attachment" "example" {
  transit_gateway_id = awscc_ec2_transit_gateway.example.id
  vpc_id             = awscc_ec2_vpc.example.id
  subnet_ids         = [awscc_ec2_subnet.example.id]

  options = {
    dns_support                        = "enable"
    ipv_6_support                      = "disable"
    appliance_mode_support             = "disable"
    security_group_referencing_support = "enable"
  }

  tags = [{
    key   = "Name"
    value = "Example-TGW-Attachment"
    }, {
    key   = "Modified By"
    value = "AWSCC"
  }]
}