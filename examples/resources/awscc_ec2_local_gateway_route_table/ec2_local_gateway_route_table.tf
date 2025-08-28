# Query for available local gateways (available in AWS Outposts)
data "aws_ec2_local_gateways" "available" {}

locals {
  local_gateway_id = length(data.aws_ec2_local_gateways.available.ids) > 0 ? data.aws_ec2_local_gateways.available.ids[0] : null
}

# Define a Local Gateway Route Table resource
resource "awscc_ec2_local_gateway_route_table" "example" {
  count = local.local_gateway_id != null ? 1 : 0

  local_gateway_id = local.local_gateway_id

  tags = [
    {
      key   = "Name"
      value = "example-route-table"
    },
    {
      key   = "Environment"
      value = "test"
    }
  ]
}
