resource "awscc_apigateway_stage" "example" {
  rest_api_id   = awscc_apigateway_rest_api.example.id
  stage_name    = "Dev"
  description   = "Dev Stage"
  deployment_id = awscc_apigateway_deployment.example.deployment_id

  variables = {
    Stack = "Dev"
  }

  method_settings = [
    {
      http_method        = "GET"
      resource_path      = "/"
      metrics_enabled    = true
      data_trace_enabled = false
    },
    {
      http_method            = "GET"
      resource_path          = "/example"
      metrics_enabled        = true
      data_trace_enabled     = false
      throttling_burst_limit = 555
    }
  ]
}

resource "awscc_apigateway_deployment" "example" {
  description = "Example deployment"
  rest_api_id = awscc_apigateway_rest_api.example.id
  depends_on  = [awscc_apigateway_method.example]
}

resource "awscc_apigateway_method" "example" {
  authorization_type = "NONE"
  http_method        = "GET"
  resource_id        = awscc_apigateway_resource.example.resource_id
  rest_api_id        = awscc_apigateway_rest_api.example.id
  integration = {
    type                    = "HTTP_PROXY"
    integration_http_method = "GET"
    uri                     = "https://ip-ranges.amazonaws.com/ip-ranges.json"
  }
}

resource "awscc_apigateway_resource" "example" {
  rest_api_id = awscc_apigateway_rest_api.example.id
  parent_id   = awscc_apigateway_rest_api.example.root_resource_id
  path_part   = "example"
}

resource "awscc_apigateway_rest_api" "example" {
  name = "example"
  endpoint_configuration = {
    types = ["REGIONAL"]
  }
}
