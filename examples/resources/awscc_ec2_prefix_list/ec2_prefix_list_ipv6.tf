resource "awscc_ec2_prefix_list" "example_ipv6_prefix_list" {
  address_family   = "IPv6"
  max_entries      = 5
  prefix_list_name = "example-ipv6-prefix-list"
  entries = [
    {
      cidr        = "2001:db8::/32"
      description = "example network"
    },
    {
      cidr        = "2001:db8:abcd:0012::0/128"
      description = "example host"
    }
  ]
}