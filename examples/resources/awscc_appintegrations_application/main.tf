# Example App Integrations Application
resource "awscc_appintegrations_application" "example" {
  name        = "example-app"
  description = "Example application created with AWSCC"
  namespace   = "contoso"

  application_source_config = {
    external_url_config = {
      access_url       = "https://example.com/app"
      approved_origins = ["https://example.com"]
    }
  }

  permissions = ["ViewQuickConnect"]

  tags = [{
    key   = "ModifiedBy"
    value = "AWSCC"
  }]
}