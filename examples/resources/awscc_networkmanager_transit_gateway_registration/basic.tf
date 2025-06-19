data "aws_partition" "current" {}
# Note: Using data.aws_region.current.region (AWS provider v6.0+)
# For AWS provider < v6.0, use data.aws_region.current.name instead
data "aws_region" "current" {}
data "aws_caller_identity" "current" {}

resource "awscc_networkmanager_global_network" "example" {}

resource "awscc_ec2_transit_gateway" "example" {}

resource "awscc_networkmanager_transit_gateway_registration" "example" {
  global_network_id   = awscc_networkmanager_global_network.example.id
  transit_gateway_arn = "arn:${data.aws_partition.current.partition}:ec2:${data.aws_region.current.region}:${data.aws_caller_identity.current.account_id}:transit-gateway/${awscc_ec2_transit_gateway.example.id}"
}