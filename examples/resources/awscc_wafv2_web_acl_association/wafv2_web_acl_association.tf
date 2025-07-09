resource "awscc_wafv2_web_acl_association" "example" {

  resource_arn = "arn:aws:apigateway:${data.aws_region.current.name}::/restapis/${awscc_apigateway_rest_api.example.id}/stages/${awscc_apigateway_stage.example.stage_name}"
  web_acl_arn  = aws_wafv2_web_acl.example.arn
}

resource "aws_wafv2_web_acl" "example" {
  name  = "web-acl-association-example"
  scope = "REGIONAL"

  default_action {
    allow {}
  }

  visibility_config {
    cloudwatch_metrics_enabled = false
    metric_name                = "friendly-metric-name"
    sampled_requests_enabled   = false
  }
}

data "aws_region" "current" {}

resource "awscc_apigateway_rest_api" "example" {
  body = jsonencode({
    openapi = "3.0.1"
    info = {
      title   = "example"
      version = "1.0"
    }
    paths = {
      "/path1" = {
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
}

resource "awscc_apigateway_deployment" "example" {
  description = "Test Apigateway Deployment"
  rest_api_id = awscc_apigateway_rest_api.example.id

  depends_on = [awscc_apigateway_rest_api.example]

}

resource "awscc_apigateway_stage" "example" {
  deployment_id = awscc_apigateway_deployment.example.deployment_id
  rest_api_id   = awscc_apigateway_rest_api.example.id
  stage_name    = "example"

  depends_on = [awscc_apigateway_deployment.example]
}