resource "awscc_ec2_ipam" "main" {
  operating_regions = [{
    region_name = "us-east-1",
  }]
}

resource "awscc_ec2_ipam_pool" "main" {
  address_family = "ipv4"
  ipam_scope_id  = awscc_ec2_ipam.main.private_default_scope_id
  locale         = "us-east-1"
}

resource "awscc_ec2_ipam_pool_cidr" "main" {
  ipam_pool_id = awscc_ec2_ipam_pool.main.id
  cidr         = "10.0.0.0/16"
}

resource "awscc_ec2_vpc" "main" {
  ipv_4_ipam_pool_id   = awscc_ec2_ipam_pool.main.id
  ipv_4_netmask_length = 20
  depends_on = [
    awscc_ec2_ipam_pool_cidr.main
  ]
}
