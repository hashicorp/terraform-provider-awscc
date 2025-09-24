resource "awscc_globalaccelerator_listener" "example" {
  accelerator_arn = awscc_globalaccelerator_accelerator.example.id
  protocol        = "TCP"

  port_ranges = [{
    from_port = "80"
    to_port   = "80"
  }]
}