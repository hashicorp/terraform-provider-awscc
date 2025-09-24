data "aws_partition" "current" {}
data "aws_region" "current" {}
data "aws_caller_identity" "current" {}

resource "awscc_networkmanager_global_network" "example" {}

resource "awscc_ec2_transit_gateway" "example" {}

data "aws_networkmanager_core_network_policy_document" "example" {
  core_network_configuration {
    asn_ranges = ["65022-65534"]

    edge_locations {
      location = data.aws_region.current.name # Core Network must have an edge location where the Transit Gateway is created
    }
  }

  segments {
    name = "segment"
  }
}

resource "awscc_networkmanager_core_network" "example" {
  global_network_id = awscc_networkmanager_global_network.example.id
  description       = "example"

  policy_document = data.aws_networkmanager_core_network_policy_document.example.json
}

resource "awscc_networkmanager_transit_gateway_peering" "example" {
  core_network_id     = awscc_networkmanager_core_network.example.id
  transit_gateway_arn = "arn:${data.aws_partition.current.partition}:ec2:${data.aws_region.current.name}:${data.aws_caller_identity.current.account_id}:transit-gateway/${awscc_ec2_transit_gateway.example.id}"
}