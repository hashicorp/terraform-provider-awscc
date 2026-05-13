resource "awscc_cases_domain" "example" {
  name = "example-cases-domain"

  tags = [
    {
      key   = "Name"
      value = "example-cases-domain"
    },
    {
      key   = "Environment"
      value = "example"
    }
  ]
}

resource "awscc_cases_field" "example" {
  domain_id   = awscc_cases_domain.example.domain_id
  name        = "example-priority-level"
  type        = "Text"
  description = "Customer priority level for case escalation and routing decisions"

  attributes = {
    text = {
      is_multiline = false
    }
  }

  tags = [
    {
      key   = "Name"
      value = "example-priority-level"
    },
    {
      key   = "Environment" 
      value = "example"
    }
  ]
}
