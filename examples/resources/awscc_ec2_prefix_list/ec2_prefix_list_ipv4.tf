resource "awscc_ec2_prefix_list" "example_prefix_list" {
  address_family   = "IPv4"
  max_entries      = 5
  prefix_list_name = "example-ipv4-prefix-list"
  entries = [
    {
      cidr        = "10.10.0.0/16"
      description = "example network"
    },
    {
      cidr        = "192.168.2.8/32"
      description = "example host"
    }
  ]
}