# Create a Connect instance first
resource "awscc_connect_instance" "example" {
  identity_management_type = "CONNECT_MANAGED"
  instance_alias           = "example-instance-${formatdate("YYYYMMDD-hhmmss", timestamp())}"
  attributes = {
    inbound_calls  = true
    outbound_calls = true
  }
  tags = [{
    key   = "Modified By"
    value = "AWSCC"
  }]
}

# Create the traffic distribution group
resource "awscc_connect_traffic_distribution_group" "example" {
  name         = "example-traffic-group"
  instance_arn = awscc_connect_instance.example.arn
  description  = "Example traffic distribution group for Connect instance"
  tags = [{
    key   = "Modified By"
    value = "AWSCC"
  }]
}