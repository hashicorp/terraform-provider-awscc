# Core Network Prefix List Association Example
# Using awscc_networkmanager_global_network and awscc_networkmanager_core_network

resource "awscc_networkmanager_global_network" "example" {
  tags = [{
    key   = "Name"
    value = "example"
    }, {
    key   = "Environment"
    value = "test"
  }]
}

data "aws_networkmanager_core_network_policy_document" "example" {
  core_network_configuration {
    asn_ranges = ["65022-65534"]

    edge_locations {
      location = "us-west-2"
    }
  }

  segments {
    name = "examplesegment"
  }
}

resource "awscc_networkmanager_core_network" "example" {
  global_network_id = awscc_networkmanager_global_network.example.id
  description       = "Example core network for prefix list association"
  policy_document   = data.aws_networkmanager_core_network_policy_document.example.json

  tags = [{
    key   = "Name"
    value = "example-core-network"
    }, {
    key   = "Environment"
    value = "test"
  }]
}

resource "aws_ec2_managed_prefix_list" "example" {
  name           = "example-prefix-list"
  address_family = "IPv4"
  max_entries    = 5

  entry {
    cidr        = "10.0.0.0/16"
    description = "Example CIDR block"
  }

  tags = {
    Name        = "example-prefix-list"
    Environment = "test"
  }
}

# Main resource: Core Network Prefix List Association
resource "awscc_networkmanager_core_network_prefix_list_association" "example" {
  core_network_id   = awscc_networkmanager_core_network.example.core_network_id
  prefix_list_arn   = aws_ec2_managed_prefix_list.example.arn
  prefix_list_alias = "examplealias"
}
