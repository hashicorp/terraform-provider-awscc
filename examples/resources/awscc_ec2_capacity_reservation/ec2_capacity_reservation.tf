resource "awscc_ec2_capacity_reservation" "example-capacity-reservation" {
  instance_type     = "t3.micro"
  instance_platform = "Linux/UNIX"
  availability_zone = "us-west-2a"
  instance_count    = 1
} 