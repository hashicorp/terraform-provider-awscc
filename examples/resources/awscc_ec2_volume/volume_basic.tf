resource "awscc_ec2_volume" "example" {
  availability_zone = "us-west-2a"
  size              = 40

  tags = [
    {
      key   = "Modified By"
      value = "AWSCC"
    }
  ]
}