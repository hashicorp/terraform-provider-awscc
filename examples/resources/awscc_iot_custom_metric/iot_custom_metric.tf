resource "awscc_iot_custom_metric" "example" {
  metric_name  = "batteryPercentage"
  metric_type  = "number"
  display_name = "Remaining battery percentage"

  tags = [{
    key   = "Modified By"
    value = "AWSCC"
  }]
}
