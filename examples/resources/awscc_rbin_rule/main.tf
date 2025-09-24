# This creates a recycle bin rule for EBS snapshots
resource "awscc_rbin_rule" "example" {
  description   = "Retain EBS snapshots for 7 days"
  resource_type = "EBS_SNAPSHOT"
  retention_period = {
    retention_period_value = 7
    retention_period_unit  = "DAYS"
  }
  resource_tags = [{
    resource_tag_key   = "Environment"
    resource_tag_value = "Production"
  }]
  tags = [{
    key   = "Modified By"
    value = "AWSCC"
  }]
}