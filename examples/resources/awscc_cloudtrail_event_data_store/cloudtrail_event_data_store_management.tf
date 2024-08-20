resource "awscc_cloudtrail_event_data_store" "example" {
  name                           = "example"
  retention_period               = 90
  federation_enabled             = false
  termination_protection_enabled = false
  multi_region_enabled           = false
  ingestion_enabled              = true
  billing_mode                   = "EXTENDABLE_RETENTION_PRICING"
  organization_enabled           = false
  advanced_event_selectors = [{
    field_selectors = [{
      field  = "eventCategory"
      equals = ["Management"]
    }]
  }]

  tags = [{
    key   = "Modified By"
    value = "AWSCC"
  }]
}
