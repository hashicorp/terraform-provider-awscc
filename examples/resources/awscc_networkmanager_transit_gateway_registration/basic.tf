data "aws_partition" "current" {}
data "aws_region" "current" {}
data "aws_caller_identity" "current" {}

resource "awscc_networkmanager_global_network" "example" {}

resource "awscc_ec2_transit_gateway" "example" {}

resource "awscc_networkmanager_transit_gateway_registration" "example" {
  global_network_id   = awscc_networkmanager_global_network.example.id
  transit_gateway_arn = "arn:${data.aws_partition.current.partition}:ec2:${data.aws_region.current.name}:${data.aws_caller_identity.current.account_id}:transit-gateway/${awscc_ec2_transit_gateway.example.id}"
}