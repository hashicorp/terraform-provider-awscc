resource "awscc_inspectorv2_filter" "example" {
  name          = "example"
  description   = "Suppression rule filtered on resource tags"
  filter_action = "NONE"
  filter_criteria = {
    resource_tags = [{
      comparison = "EQUALS"
      key        = "Modified BY"
      value      = "AWSCC"
      }
    ]
  }

}