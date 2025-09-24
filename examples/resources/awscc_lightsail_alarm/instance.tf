resource "awscc_lightsail_instance" "example" {
  blueprint_id  = "amazon_linux_2023"
  bundle_id     = "nano_3_0"
  instance_name = "example-instance"
}

resource "awscc_lightsail_alarm" "example" {
  alarm_name              = "example-alarm"
  comparison_operator     = "GreaterThanThreshold"
  evaluation_periods      = 1
  metric_name             = "CPUUtilization"
  monitored_resource_name = awscc_lightsail_instance.example.instance_name
  threshold               = 90
}
