resource "awscc_lightsail_instance" "example" {
  blueprint_id  = "nginx"
  bundle_id     = "nano_3_0"
  instance_name = "example-instance"
}

resource "awscc_lightsail_load_balancer" "example" {
  instance_port      = 80
  load_balancer_name = "example-lb"
  attached_instances = [awscc_lightsail_instance.example.instance_name]
}

# Since there is no resource for contact method, we need to create it using null_resource and the AWS CLI instead
resource "null_resource" "example_contact_method" {
  provisioner "local-exec" {
    command = "aws lightsail create-contact-method --protocol Email --contact-endpoint admin@example.com"
  }
  provisioner "local-exec" {
    when    = destroy
    command = "aws lightsail delete-contact-method --protocol Email"
  }
}

resource "awscc_lightsail_alarm" "example" {
  alarm_name              = "example-alarm"
  comparison_operator     = "GreaterThanOrEqualToThreshold"
  evaluation_periods      = 1
  metric_name             = "UnhealthyHostCount"
  monitored_resource_name = awscc_lightsail_load_balancer.example.load_balancer_name
  threshold               = 1
  contact_protocols       = ["Email"]
  notification_enabled    = true
  notification_triggers   = ["ALARM"]
  treat_missing_data      = "ignore"
  depends_on              = [null_resource.example_contact_method]
}
