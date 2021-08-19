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

# resource "awscc_ec2_dhcp_options" "test" {
#   domain_name          = "service.tf"
#   domain_name_servers  = ["127.0.0.1", "10.0.0.2"]
#   ntp_servers          = ["127.0.0.1"]
#   netbios_name_servers = ["127.0.0.1"]
#   netbios_node_type    = 2

#   tags = [
#     {
#       key   = "Name"
#       value = "AWS CC testing"
#     }
#   ]
# }

# resource "awscc_athena_work_group" "test" {
#   name = "awscc-testing"

#   work_group_configuration = {
#     bytes_scanned_cutoff_per_query = 12582912

#     result_configuration = {
#       encryption_configuration = {
#         encryption_option = "SSE_S3"
#       }
#     }
#   }
# }

resource "awscc_xray_group" "test" {
  filter_expression = "responsetime > 5"
  group_name        = "ewbankkit-cc-test"
}