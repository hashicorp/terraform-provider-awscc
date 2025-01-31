# Create cost anomaly monitor for AWS cost
resource "awscc_ce_anomaly_monitor" "cost_monitor" {
  monitor_name = "aws-cost-monitor"
  monitor_type = "DIMENSIONAL"
  monitor_specification = jsonencode({
    dimensions = {
      key = "SERVICE"
    }
  })

  tags = [{
    key   = "Modified By"
    value = "AWSCC"
  }]
}
