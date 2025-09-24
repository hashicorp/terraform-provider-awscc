resource "awscc_apigateway_rest_api" "DemoRestAPI" {
  name = "DemoRestAPI"
  endpoint_configuration = {
    types = ["REGIONAL"]
  }
  body = jsonencode({
    openapi = "3.0.1"
    info = {
      title   = "DemoRestAPI"
      version = "1.0"
    }
    paths = {
      "/path1" = {
        get = {
          x-amazon-apigateway-integration = {
            payloadFormatVersion = "1.0"
            httpMethod           = "GET"
            type                 = "HTTP_PROXY"
            uri                  = "https://ip-ranges.amazonaws.com/ip-ranges.json"
          }
        }
      }
    }
  })
}

resource "awscc_apigateway_resource" "DemoAPIGatewayResource" {
  rest_api_id = awscc_apigateway_rest_api.DemoRestAPI.id
  parent_id   = awscc_apigateway_rest_api.DemoRestAPI.root_resource_id
  path_part   = "DemoAPIGatewayResource"
}

