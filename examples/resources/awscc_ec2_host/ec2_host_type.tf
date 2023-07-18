resource "awscc_ec2_host" "dedicated_host1" {
  availability_zone = "us-east-1b"
  auto_placement    = "on"
  instance_type     = "m5.xlarge"
}