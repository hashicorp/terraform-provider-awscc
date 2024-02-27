resource "awscc_ec2_capacity_reservation_fleet" "example-reserved-fleet" {
  total_target_capacity   = 480
  allocation_strategy     = "prioritized"
  instance_match_criteria = "open"
  tenancy                 = "default"
  end_date                = "2023-12-01T23:59:59Z"

  instance_type_specifications = [{
    instance_type     = "m5.4xlarge"
    instance_platform = "Windows"
    weight            = 16
    availability_zone = "us-west-2a"
    ebs_optimized     = true
    priority          = 2
    },
    {
      instance_type     = "m5.12xlarge"
      instance_platform = "Windows"
      weight            = 48
      availability_zone = "us-west-2a"
      ebs_optimized     = true
      priority          = 1
  }]
}