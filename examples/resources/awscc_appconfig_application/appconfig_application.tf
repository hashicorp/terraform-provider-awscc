resource "awscc_appconfig_application" "example" {
  name        = "example"
  description = "Example application"

  tags = [{
    key   = "ModifiedBy"
    value = "AWSCC"
  }]
}

