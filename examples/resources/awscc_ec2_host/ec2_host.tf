resource "awscc_ec2_host" "this" {
  instance_type     = "c5.large"
  availability_zone = "us-east-2a"
  host_recovery     = "on"
  auto_placement    = "on"
}