data "aws_caller_identity" "current" {}
data "aws_region" "current" {}

locals {
  account_id = data.aws_caller_identity.current.account_id
  region     = data.aws_region.current.name
}

resource "awscc_networkmanager_global_network" "example" {
  description = "Example Global Network"
  tags = [{
    key   = "Modified By"
    value = "AWSCC"
  }]
}

# Core Network - segment-actions cannot reference attachment IDs during creation
# as attachments don't exist yet. This creates circular dependencies on both
# create and destroy operations. Use blackhole or add segment-actions later.
resource "awscc_networkmanager_core_network" "example" {
  description       = "Example Core Network"
  global_network_id = awscc_networkmanager_global_network.example.id
  policy_document = jsonencode({
    "version" : "2021.12",
    "core-network-configuration" : {
      "vpn-ecmp-support" : true,
      "asn-ranges" : ["64512-65534"],
      "edge-locations" : [{
        "location" : local.region
      }]
    },
    "segments" : [{
      "name" : "shared",
      "description" : "Segment for shared services",
      "require-attachment-acceptance" : false
    }]
  })
  tags = [{
    key   = "Modified By"
    value = "AWSCC"
  }]
}

resource "awscc_ec2_vpc" "example" {
  cidr_block = "10.0.0.0/16"
  tags = [{
    key   = "Name"
    value = "example-vpc"
  }]
}

resource "awscc_ec2_subnet" "example_subnet1" {
  vpc_id            = awscc_ec2_vpc.example.id
  cidr_block        = "10.0.1.0/24"
  availability_zone = "${local.region}a"
  tags = [{
    key   = "Name"
    value = "example-subnet-1"
  }]
}

resource "awscc_ec2_subnet" "example_subnet2" {
  vpc_id            = awscc_ec2_vpc.example.id
  cidr_block        = "10.0.2.0/24"
  availability_zone = "${local.region}b"
  tags = [{
    key   = "Name"
    value = "example-subnet-2"
  }]
}

resource "awscc_networkmanager_vpc_attachment" "example" {
  core_network_id = awscc_networkmanager_core_network.example.id
  vpc_arn         = format("arn:aws:ec2:%s:%s:vpc/%s", local.region, local.account_id, awscc_ec2_vpc.example.id)
  subnet_arns = [
    format("arn:aws:ec2:%s:%s:subnet/%s", local.region, local.account_id, awscc_ec2_subnet.example_subnet1.id),
    format("arn:aws:ec2:%s:%s:subnet/%s", local.region, local.account_id, awscc_ec2_subnet.example_subnet2.id)
  ]
  options = {
    appliance_mode_support = false
    ipv_6_support          = false
  }
  proposed_segment_change = {
    segment_name = "shared"
  }
  tags = [{
    key   = "Modified By"
    value = "AWSCC"
  }]
}
