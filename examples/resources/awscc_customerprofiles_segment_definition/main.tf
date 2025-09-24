# Create a Customer Profiles Domain first (required for the segment definition)
resource "awscc_customerprofiles_domain" "example" {
  domain_name             = "example-domain"
  default_expiration_days = 365
  tags = [{
    key   = "ModifiedBy"
    value = "AWSCC"
  }]
}

# Create the segment definition
resource "awscc_customerprofiles_segment_definition" "example" {
  domain_name             = awscc_customerprofiles_domain.example.domain_name
  segment_definition_name = "example-segment"
  display_name            = "Example Segment"
  description             = "Example segment definition to demonstrate configuration"

  segment_groups = {
    include = "ALL"
    groups = [
      {
        source_type = "ALL"
        dimensions = [
          {
            profile_attributes = {
              first_name = {
                dimension_type = "INCLUSIVE"
                values         = ["John"]
              }
              personal_email_address = {
                dimension_type = "INCLUSIVE"
                values         = ["john@example.com"]
              }
            }
          }
        ]
      }
    ]
  }

  tags = [{
    key   = "ModifiedBy"
    value = "AWSCC"
  }]
}