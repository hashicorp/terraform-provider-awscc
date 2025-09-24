resource "awscc_apigateway_documentation_part" "query" {
  location = {
    type   = "QUERY_PARAMETER"
    method = "GET"
    path   = "/example"
    name   = "limit"
  }
  properties = jsonencode({
    "description" : "Parameter to control max number of records to return"
  })
  rest_api_id = awscc_apigateway_rest_api.example.id
}

resource "awscc_apigateway_rest_api" "example" {
  name = "example_api"
}
