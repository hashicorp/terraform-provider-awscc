# EC2 Traffic Mirror Filter resource
resource "awscc_ec2_traffic_mirror_filter" "example" {
  description      = "Example traffic mirror filter"
  network_services = ["amazon-dns"]

  tags = [
    {
      key   = "Name"
      value = "example-filter"
    },
    {
      key   = "Environment"
      value = "example"
    }
  ]
}
