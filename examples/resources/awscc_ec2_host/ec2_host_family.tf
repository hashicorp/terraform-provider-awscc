resource "awscc_ec2_host" "dedicated_host" {
  availability_zone = "us-east-1a"
  auto_placement    = "on"
  instance_family   = "m5"
}