resource "awscc_apigateway_model" "example" {
  rest_api_id  = awscc_apigateway_rest_api.example.id
  name         = "example"
  description  = "example model"
  content_type = "application/json"

  schema = jsonencode({
    type = "object"
  })
}

resource "awscc_apigateway_rest_api" "example" {
  name        = "exampleAPI"
  description = "Example API"
}