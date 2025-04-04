# Create a cluster first
resource "awscc_route53recoverycontrol_cluster" "example" {
  name = "example-cluster"
  tags = [{
    key   = "Modified By"
    value = "AWSCC"
  }]
}

# Create a control panel
resource "awscc_route53recoverycontrol_control_panel" "example" {
  name        = "example-control-panel"
  cluster_arn = awscc_route53recoverycontrol_cluster.example.cluster_arn
  tags = [{
    key   = "Modified By"
    value = "AWSCC"
  }]
}

# Create routing controls
resource "awscc_route53recoverycontrol_routing_control" "region1" {
  name              = "region1-control"
  cluster_arn       = awscc_route53recoverycontrol_cluster.example.cluster_arn
  control_panel_arn = awscc_route53recoverycontrol_control_panel.example.control_panel_arn
}

resource "awscc_route53recoverycontrol_routing_control" "region2" {
  name              = "region2-control"
  cluster_arn       = awscc_route53recoverycontrol_cluster.example.cluster_arn
  control_panel_arn = awscc_route53recoverycontrol_control_panel.example.control_panel_arn
}

# Create a safety rule (assertion rule type)
resource "awscc_route53recoverycontrol_safety_rule" "example" {
  name              = "example-safety-rule"
  control_panel_arn = awscc_route53recoverycontrol_control_panel.example.control_panel_arn

  assertion_rule = {
    asserted_controls = [
      awscc_route53recoverycontrol_routing_control.region1.routing_control_arn,
      awscc_route53recoverycontrol_routing_control.region2.routing_control_arn
    ]
    wait_period_ms = 5000
  }

  rule_config = {
    inverted  = false
    threshold = 1
    type      = "ATLEAST"
  }

  tags = [{
    key   = "Modified By"
    value = "AWSCC"
  }]
}