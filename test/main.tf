terraform {
  required_providers {

    aws = {
      source = "hashicorp/aws"
    }

    awscc = {
      source = "hashicorp/awscc"
    }
  }
}

provider "aws" {}
provider "awscc" {}

resource "awscc_ec2_dhcp_options" "test" {
  domain_name          = "service.tf"
  domain_name_servers  = ["127.0.0.1", "10.0.0.2"]
  ntp_servers          = ["127.0.0.1"]
  netbios_name_servers = ["127.0.0.1"]
  netbios_node_type    = 2

  tags = [
    {
      key   = "Name"
      value = "AWS CC testing"
    }
  ]
}
