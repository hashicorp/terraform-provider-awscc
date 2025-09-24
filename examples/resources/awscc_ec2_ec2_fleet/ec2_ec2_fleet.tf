resource "awscc_ec2_ec2_fleet" "example_fleet" {
  launch_template_configs = [{
    launch_template_specification = {
      launch_template_id = awscc_ec2_launch_template.example_launch_template.id
      version            = "$Latest"
    }
  }]
  target_capacity_specification = {
    default_target_capacity_type = "spot"
    total_target_capacity        = 5
  }
}