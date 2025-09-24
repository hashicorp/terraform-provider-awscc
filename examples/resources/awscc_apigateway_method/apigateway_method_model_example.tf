
resource "awscc_apigateway_rest_api" "terraform_apigateway_rest_api" {
  name = "TestRestApi"
  endpoint_configuration = {
    types = [
      "REGIONAL"
    ]
  }
}

resource "awscc_apigateway_method" "terraform_apigateway_method" {
  http_method        = "GET"
  authorization_type = "NONE"
  integration = {
    type = "MOCK"
  }
  rest_api_id = awscc_apigateway_rest_api.terraform_apigateway_rest_api.id
  resource_id = awscc_apigateway_rest_api.terraform_apigateway_rest_api.root_resource_id

  method_responses = [{ status_code = "200", response_models = { "application/json" : "Empty" }, response_parameters = { "method.response.header.Content-Type" = false } }]

  depends_on = [awscc_apigateway_rest_api.terraform_apigateway_rest_api]
}