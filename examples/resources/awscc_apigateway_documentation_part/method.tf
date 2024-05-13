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
