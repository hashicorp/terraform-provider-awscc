# Create Customer Profiles Domain
resource "awscc_customerprofiles_domain" "example" {
  domain_name = "example-domain"

  default_expiration_days = 365

  tags = [{
    key   = "ModifiedBy"
    value = "AWSCC"
  }]
}

# Create Customer Profiles Object Type
resource "awscc_customerprofiles_object_type" "example" {
  domain_name      = awscc_customerprofiles_domain.example.domain_name
  object_type_name = "example-object-type"
  description      = "Example Customer Profiles Object Type"

  allow_profile_creation = true
  expiration_days        = 365

  # Define fields for the object type
  fields = [
    {
      name = "FirstName"
      object_type_field = {
        content_type = "STRING"
        source       = "_source.FirstName"
        target       = "_profile.FirstName"
      }
    },
    {
      name = "LastName"
      object_type_field = {
        content_type = "STRING"
        source       = "_source.LastName"
        target       = "_profile.LastName"
      }
    },
    {
      name = "EmailAddress"
      object_type_field = {
        content_type = "EMAIL_ADDRESS"
        source       = "_source.EmailAddress"
        target       = "_profile.EmailAddress"
      }
    }
  ]

  # Define keys for the object type
  keys = [
    {
      name = "email"
      object_type_key_list = [
        {
          field_names          = ["EmailAddress"]
          standard_identifiers = ["PROFILE", "UNIQUE"]
        }
      ]
    }
  ]

  tags = [{
    key   = "ModifiedBy"
    value = "AWSCC"
  }]
}