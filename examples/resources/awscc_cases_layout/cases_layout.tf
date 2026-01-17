resource "awscc_cases_layout" "example" {
  domain_id = awscc_cases_domain.example.domain_id
  name      = "example-layout"
  
  content = {
    basic = {
      top_panel = {
        sections = [{
          field_group = {
            name = "Summary"
            fields = [{
              id = awscc_cases_field.example.field_id
            }]
          }
        }]
      }
    }
  }
  
  tags = [{
    key   = "Environment"
    value = "test"
  }]
}

resource "awscc_cases_domain" "example" {
  name = "example-cases-domain-for-layout"
  
  tags = [{
    key   = "Environment"
    value = "test"
  }]
}

resource "awscc_cases_field" "example" {
  domain_id = awscc_cases_domain.example.domain_id
  name      = "example-field-for-layout"
  type      = "Text"
  
  tags = [{
    key   = "Environment"
    value = "test"
  }]
}
