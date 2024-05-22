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

  depends_on = [awscc_apigateway_rest_api.terraform_apigateway_rest_api]
}

resource "awscc_apigateway_deployment" "terraform_apigateway_deployment" {
  description = "Test Apigateway Deployment"
  rest_api_id = awscc_apigateway_rest_api.terraform_apigateway_rest_api.id
  stage_description = {
    description = "Test stage description"
  }
  stage_name = "Test Stage"

  depends_on = [awscc_apigateway_method.terraform_apigateway_method, awscc_apigateway_rest_api.terraform_apigateway_rest_api]
}

