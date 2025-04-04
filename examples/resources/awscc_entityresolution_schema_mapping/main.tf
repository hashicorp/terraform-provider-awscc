resource "awscc_entityresolution_schema_mapping" "example" {
  schema_name = "example-schema-mapping"
  description = "Example schema mapping for customer data"

  mapped_input_fields = [
    {
      field_name = "id"
      type       = "UNIQUE_ID"
    },
    {
      field_name = "name"
      type       = "NAME"
      match_key  = "NAME"
    },
    {
      field_name = "email"
      type       = "EMAIL_ADDRESS"
      match_key  = "EMAIL"
      hashed     = true
    },
    {
      field_name = "address"
      type       = "ADDRESS"
      match_key  = "ADDRESS"
      group_name = "contact"
    }
  ]

  tags = [{
    key   = "Modified By"
    value = "AWSCC"
  }]
}