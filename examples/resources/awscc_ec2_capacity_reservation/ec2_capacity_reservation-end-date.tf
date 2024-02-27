resource "awscc_ec2_capacity_reservation" "example-capacity-reservation-end-date" {
  instance_type     = "t3.micro"
  instance_platform = "Linux/UNIX"
  availability_zone = "us-west-2a"
  instance_count    = 3
  end_date_type     = "limited"
  end_date          = "2023-12-01T23:59:59Z"
} 