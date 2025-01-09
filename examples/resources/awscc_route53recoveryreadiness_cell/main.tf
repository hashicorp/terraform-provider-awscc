# Create a Route53 Recovery Readiness Cell
resource "awscc_route53recoveryreadiness_cell" "example" {
  cell_name = "my-recovery-cell"

  tags = [{
    key   = "Modified By"
    value = "AWSCC"
  }]
}