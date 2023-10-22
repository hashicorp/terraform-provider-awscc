resource "awscc_apigateway_rest_api" "example" {
  name = "ExampleAPI"
}

resource "awscc_apigateway_resource" "example" {
  rest_api_id = awscc_apigateway_rest_api.example.id
  parent_id   = awscc_apigateway_rest_api.example.root_resource_id
  path_part   = "path"
}

resource "awscc_apigateway_method" "example" {
  rest_api_id = awscc_apigateway_rest_api.example.id
  resource_id = awscc_apigateway_resource.example.resource_id
  http_method = "GET"

  authorization_type = "NONE"

  integration = {
    type = "MOCK"
  }
}