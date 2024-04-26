resource "awscc_ec2_security_group" "allow_tls" {
  group_description = "Allow TLS inbound traffic and all outbound traffic"
  vpc_id            = awscc_ec2_vpc.selected.id

  tags = [
    {
      key   = "Name"
      value = "allow_tls"
    }
  ]
}

resource "awscc_ec2_vpc_cidr_block" "selected" {
  amazon_provided_ipv_6_cidr_block = true
  vpc_id                           = awscc_ec2_vpc.selected.id
}

resource "awscc_ec2_vpc" "selected" {
  cidr_block           = "10.0.0.0/16"
  enable_dns_hostnames = true
}


resource "awscc_ec2_security_group_ingress" "allow_tls_ipv4" {
  group_id    = awscc_ec2_security_group.allow_tls.id
  cidr_ip     = awscc_ec2_vpc.selected.cidr_block
  from_port   = 443
  ip_protocol = "tcp"
  to_port     = 443
}

resource "awscc_ec2_security_group_ingress" "allow_tls_ipv6" {
  group_id    = awscc_ec2_security_group.allow_tls.id
  cidr_ipv_6  = awscc_ec2_vpc_cidr_block.selected.ipv_6_cidr_block
  from_port   = 443
  ip_protocol = "tcp"
  to_port     = 443
}

resource "awscc_ec2_security_group_egress" "allow_all_traffic_ipv4" {
  group_id    = awscc_ec2_security_group.allow_tls.id
  cidr_ip     = "0.0.0.0/0"
  ip_protocol = "-1" # semantically equivalent to all ports
}

resource "awscc_ec2_security_group_egress" "allow_all_traffic_ipv6" {
  group_id    = awscc_ec2_security_group.allow_tls.id
  cidr_ipv_6  = "::/0"
  ip_protocol = "-1" # semantically equivalent to all ports
}