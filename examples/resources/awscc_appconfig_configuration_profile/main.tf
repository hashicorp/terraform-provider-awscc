# AppConfig Application is required before creating a configuration profile
resource "awscc_appconfig_application" "example" {
  name        = "example-app"
  description = "Example AppConfig application"

  tags = [{
    key   = "Modified By"
    value = "AWSCC"
  }]
}

# Configuration profile with feature flags type
resource "awscc_appconfig_configuration_profile" "feature_flags" {
  name           = "example-feature-flags"
  application_id = awscc_appconfig_application.example.id
  location_uri   = "hosted"
  description    = "Example feature flags configuration profile"
  type           = "AWS.AppConfig.FeatureFlags"

  validators = [{
    type = "JSON_SCHEMA"
    content = jsonencode({
      "$schema" = "http://json-schema.org/draft-04/schema#"
      type      = "object"
      properties = {
        flags = {
          type = "object"
        }
      }
      required = ["flags"]
    })
  }]

  tags = [{
    key   = "Modified By"
    value = "AWSCC"
  }]
}

# Configuration profile with freeform type
resource "awscc_appconfig_configuration_profile" "freeform" {
  name           = "example-freeform"
  application_id = awscc_appconfig_application.example.id
  location_uri   = "hosted"
  description    = "Example freeform configuration profile"
  type           = "AWS.Freeform"

  validators = [{
    type = "JSON_SCHEMA"
    content = jsonencode({
      "$schema"   = "http://json-schema.org/draft-04/schema#"
      type        = "object"
      description = "Basic schema validation"
      required    = ["key1"]
      properties = {
        key1 = {
          type = "string"
        }
      }
    })
  }]

  tags = [{
    key   = "Modified By"
    value = "AWSCC"
  }]
}