resource "awscc_cases_domain" "example" {
  name = "example-cases-domain-for-rule"

  tags = [{
    key   = "Environment"
    value = "test"
  }]
}

resource "awscc_cases_field" "example" {
  domain_id = awscc_cases_domain.example.domain_id
  name      = "example-field-for-rule"
  type      = "Text"

  tags = [{
    key   = "Environment"
    value = "test"
  }]
}

resource "awscc_cases_case_rule" "example" {
  domain_id = awscc_cases_domain.example.domain_id
  name      = "example-case-rule"

  rule = {
    required = {
      default_value = false
      conditions = [{
        equal_to = {
          operand_one = {
            field_id = awscc_cases_field.example.field_id
          }
          operand_two = {
            string_value = "test"
          }
          result = true
        }
      }]
    }
  }

  tags = [{
    key   = "Environment"
    value = "test"
  }]
}
