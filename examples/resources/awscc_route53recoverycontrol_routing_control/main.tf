# First create a cluster
resource "awscc_route53recoverycontrol_cluster" "example" {
  name = "example-cluster"
  tags = [{
    key   = "Modified By"
    value = "AWSCC"
  }]
}

# Then create a control panel
resource "awscc_route53recoverycontrol_control_panel" "example" {
  name        = "example-control-panel"
  cluster_arn = awscc_route53recoverycontrol_cluster.example.cluster_arn
  tags = [{
    key   = "Modified By"
    value = "AWSCC"
  }]
}

# Finally create the routing control
resource "awscc_route53recoverycontrol_routing_control" "example" {
  name              = "example-routing-control"
  cluster_arn       = awscc_route53recoverycontrol_cluster.example.cluster_arn
  control_panel_arn = awscc_route53recoverycontrol_control_panel.example.control_panel_arn
}