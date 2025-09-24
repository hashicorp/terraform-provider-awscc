resource "awscc_apigateway_documentation_version" "example" {
  documentation_version = "v1"
  rest_api_id           = awscc_apigateway_rest_api.example.id
  description           = "API documentation version snapshot v1"
  depends_on            = [awscc_apigateway_documentation_part.example]
}

resource "awscc_apigateway_documentation_part" "example" {
  location = {
    type   = "METHOD"
    method = "GET"
    path   = "/example"
  }
  properties = jsonencode({
    "description" : "Example description"
  })
  rest_api_id = awscc_apigateway_rest_api.example.id
}

resource "awscc_apigateway_rest_api" "example" {
  name = "example_api"
}
