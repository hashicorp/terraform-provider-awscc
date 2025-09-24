resource "awscc_ec2_network_insights_access_scope_analysis" "example" {
  network_insights_access_scope_id = awscc_ec2_network_insights_access_scope.example.id

  tags = [{
    key   = "Name"
    value = "analysis1"
  }]
}