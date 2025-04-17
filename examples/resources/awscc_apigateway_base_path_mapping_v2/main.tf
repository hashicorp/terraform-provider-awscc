# Create ACM certificate for the domain
resource "aws_acm_certificate" "example" {
  domain_name       = "example.com"
  validation_method = "DNS"

  tags = {
    Environment = "test"
  }

  lifecycle {
    create_before_destroy = true
  }
}

# API Gateway REST API
resource "awscc_apigateway_rest_api" "example" {
  name = "example-api"
  tags = [{
    key   = "Modified By"
    value = "AWSCC"
  }]
}

# Sample API Gateway Method and Integration using AWS provider
resource "aws_api_gateway_method" "example" {
  rest_api_id   = awscc_apigateway_rest_api.example.id
  resource_id   = awscc_apigateway_rest_api.example.root_resource_id
  http_method   = "GET"
  authorization = "NONE"
}

resource "aws_api_gateway_integration" "example" {
  rest_api_id = awscc_apigateway_rest_api.example.id
  resource_id = awscc_apigateway_rest_api.example.root_resource_id
  http_method = aws_api_gateway_method.example.http_method
  type        = "MOCK"

  request_templates = {
    "application/json" = jsonencode({
      statusCode = 200
    })
  }
}

# API Gateway Deployment using AWS provider
resource "aws_api_gateway_deployment" "example" {

  rest_api_id = awscc_apigateway_rest_api.example.id

  # Make sure we have method and integration before deploying
  depends_on = [
    aws_api_gateway_method.example,
    aws_api_gateway_integration.example
  ]

  lifecycle {
    create_before_destroy = true
  }
}

# API Gateway Stage
resource "awscc_apigateway_stage" "example" {
  rest_api_id   = awscc_apigateway_rest_api.example.id
  deployment_id = aws_api_gateway_deployment.example.id
  stage_name    = "prod"

  tags = [{
    key   = "Modified By"
    value = "AWSCC"
  }]
}

# API Gateway Domain Name
resource "awscc_apigateway_domain_name_v2" "example" {
  domain_name = "example.com"
  endpoint_configuration = {
    types                    = ["REGIONAL"]
    regional_certificate_arn = aws_acm_certificate.example.arn
  }

  tags = [{
    key   = "Modified By"
    value = "AWSCC"
  }]
}

# API Gateway Base Path Mapping
resource "awscc_apigateway_base_path_mapping_v2" "example" {
  domain_name_arn = awscc_apigateway_domain_name_v2.example.domain_name_arn
  rest_api_id     = awscc_apigateway_rest_api.example.id
  base_path       = "v1"
  stage           = awscc_apigateway_stage.example.stage_name
}