resource "awscc_networkmanager_global_network" "example" {}

data "aws_networkmanager_core_network_policy_document" "example" {
  core_network_configuration {
    asn_ranges = ["65022-65534"]

    edge_locations {
      location = "us-west-2"
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