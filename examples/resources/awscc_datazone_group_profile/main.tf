# Example DataZone Group Profile
resource "awscc_datazone_group_profile" "example" {
  domain_identifier = "dzd_example123"
  group_identifier  = "example-group"
  status           = "ASSIGNED"
}