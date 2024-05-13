resource "awscc_apigateway_rest_api" "example" {
  body = jsonencode({
    openapi = "3.0.1"
    info = {
      title   = "example"
      version = "1.0"
    }
    paths = {
      "/example" = {
        get = {
          x-amazon-apigateway-integration = {
            httpMethod           = "GET"
            payloadFormatVersion = "1.0"
            type                 = "HTTP_PROXY"
            uri                  = "https://ip-ranges.amazonaws.com/ip-ranges.json"
          }
        }
      }
    }
  })

  name = "example"

  endpoint_configuration = {
    types = ["REGIONAL"]
  }
}

resource "awscc_apigateway_deployment" "example" {
  description = "Example Deployment"
  rest_api_id = awscc_apigateway_rest_api.example.id
}

resource "awscc_apigateway_stage" "example" {
  rest_api_id   = awscc_apigateway_rest_api.example.id
  deployment_id = awscc_apigateway_deployment.example.deployment_id
  stage_name    = "dev"
}

