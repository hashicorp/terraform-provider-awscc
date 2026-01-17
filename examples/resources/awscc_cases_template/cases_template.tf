resource "awscc_cases_template" "example" {
  domain_id = awscc_cases_domain.example.domain_id
  name      = "example-template"
  
  layout_configuration = {
    default_layout = awscc_cases_layout.example.layout_id
  }
  
  required_fields = [{
    field_id = awscc_cases_field.example.field_id
  }]
  
  tags = [{
    key   = "Environment"
    value = "test"
  }]
}

resource "awscc_cases_domain" "example" {
  name = "example-cases-domain-for-template"
  
  tags = [{
    key   = "Environment"
    value = "test"
  }]
}

resource "awscc_cases_field" "example" {
  domain_id = awscc_cases_domain.example.domain_id
  name      = "example-field-for-template"
  type      = "Text"
  
  tags = [{
    key   = "Environment"
    value = "test"
  }]
}

resource "awscc_cases_layout" "example" {
  domain_id = awscc_cases_domain.example.domain_id
  name      = "example-layout-for-template"
  
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
