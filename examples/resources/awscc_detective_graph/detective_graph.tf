resource "awscc_detective_graph" "example" {
  auto_enable_members = true
  tags = [{
    key   = "Modified By"
    value = "AWSCC"
  }]
}
