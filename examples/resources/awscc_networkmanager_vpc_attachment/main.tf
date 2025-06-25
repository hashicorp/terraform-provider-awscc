data "aws_caller_identity" "current" {}
# Note: Using data.aws_region.current.region (AWS provider v6.0+)
# For AWS provider < v6.0, use data.aws_region.current.name instead
data "aws_region" "current" {}

# Create VPC and subnets
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
  availability_zone = "${data.aws_region.current.region}a"
  tags = [{
    key   = "Name"
    value = "example-subnet-1"
  }]
}

resource "awscc_ec2_subnet" "example_subnet2" {
  vpc_id            = awscc_ec2_vpc.example.id
  cidr_block        = "10.0.2.0/24"
  availability_zone = "${data.aws_region.current.region}b"
  tags = [{
    key   = "Name"
    value = "example-subnet-2"
  }]
}

# Create Network Manager resources
resource "awscc_networkmanager_global_network" "example" {
  description = "Example Global Network"
  tags = [{
    key   = "Modified By"
    value = "AWSCC"
  }]
}

resource "awscc_networkmanager_core_network" "example" {
  description       = "Example Core Network"
  global_network_id = awscc_networkmanager_global_network.example.id
  policy_document = jsonencode({
    "version" : "2021.12",
    "core-network-configuration" : {
      "vpn-ecmp-support" : true,
      "asn-ranges" : [
        "64512-65534"
      ],
      "edge-locations" : [
        {
          "location" : data.aws_region.current.region
        }
      ]
    },
    "segments" : [
      {
        "name" : "shared",
        "description" : "Segment for shared services",
        "require-attachment-acceptance" : false
      }
    ],
    "segment-actions" : [
      {
        "action" : "create-route",
        "destination-cidr-blocks" : [
          "0.0.0.0/0"
        ],
        "destinations" : [
          "attachment"
        ],
        "segment" : "shared"
      }
    ]
  })
  tags = [{
    key   = "Modified By"
    value = "AWSCC"
  }]
}

# Create VPC Attachment
resource "awscc_networkmanager_vpc_attachment" "example" {
  core_network_id = awscc_networkmanager_core_network.example.id
  vpc_arn         = format("arn:aws:ec2:%s:%s:vpc/%s", data.aws_region.current.region, data.aws_caller_identity.current.account_id, awscc_ec2_vpc.example.id)
  subnet_arns = [
    format("arn:aws:ec2:%s:%s:subnet/%s", data.aws_region.current.region, data.aws_caller_identity.current.account_id, awscc_ec2_subnet.example_subnet1.id),
    format("arn:aws:ec2:%s:%s:subnet/%s", data.aws_region.current.region, data.aws_caller_identity.current.account_id, awscc_ec2_subnet.example_subnet2.id)
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