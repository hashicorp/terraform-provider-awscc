# Knowledge base resource is required for the quick response
resource "awscc_wisdom_knowledge_base" "example" {
  name = "example-knowledge-base"

  # Quick response knowledge base type
  knowledge_base_type = "QUICK_RESPONSES"

  tags = [
    {
      key   = "Environment"
      value = "test"
    },
    {
      key   = "Name"
      value = "example-knowledge-base"
    }
  ]
}

# Quick response resource
resource "awscc_wisdom_quick_response" "example" {
  knowledge_base_arn = awscc_wisdom_knowledge_base.example.knowledge_base_arn
  name               = "example-quick-response"

  # Content with proper structure based on the schema
  content = {
    content = jsonencode({
      Content = {
        PlainText = {
          Text = "This is an example quick response text that can be used to answer common customer questions."
        }
      }
    })
  }

  content_type = "application/x.quickresponse;format=plain"
  description  = "Example quick response for demonstration"
  is_active    = true
  language     = "en_US"
  shortcut_key = "exqr"

  channels = ["Chat"]

  tags = [
    {
      key   = "Environment"
      value = "test"
    },
    {
      key   = "Name"
      value = "example-knowledge-base"
    }
  ]
}
