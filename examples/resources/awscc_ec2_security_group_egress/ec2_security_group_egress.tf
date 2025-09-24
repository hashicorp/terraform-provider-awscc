resource "awscc_ec2_security_group_egress" "allow_https_traffic_ipv4" {
  group_id    = awscc_ec2_security_group.example.id
  cidr_ip     = "0.0.0.0/0"
  ip_protocol = "tcp"
  from_port   = 443
  to_port     = 443
  description = "Outbound rule to allow https traffic"
}

resource "awscc_ec2_security_group" "example" {
  group_description = "Example SG"
  vpc_id            = awscc_ec2_vpc.selected.id

  tags = [
    {
      key   = "Name"
      value = "example_sg"
    }
  ]
}

resource "awscc_ec2_vpc" "selected" {
  cidr_block = "10.0.0.0/16"
}