# Create a cluster first as it's required for the control panel
resource "awscc_route53recoverycontrol_cluster" "example" {
  name = "example-cluster"
  tags = [{
    key   = "Modified By"
    value = "AWSCC"
  }]
}

# Create the control panel
resource "awscc_route53recoverycontrol_control_panel" "example" {
  name        = "example-control-panel"
  cluster_arn = awscc_route53recoverycontrol_cluster.example.cluster_arn

  tags = [{
    key   = "Modified By"
    value = "AWSCC"
  }]
}