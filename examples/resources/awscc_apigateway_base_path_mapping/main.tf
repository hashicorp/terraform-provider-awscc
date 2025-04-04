# Generate a SSL certificate for the domain
resource "aws_acm_certificate" "example" {
  domain_name       = "api.example.com"
  validation_method = "DNS"

  lifecycle {
    create_before_destroy = true
  }

  tags = {
    Environment = "test"
  }
}

# Create a REST API
resource "awscc_apigateway_rest_api" "example" {
  name = "example-api"
  tags = [{
    key   = "Modified By"
    value = "AWSCC"
  }]
}

# Create an API Stage
resource "awscc_apigateway_stage" "example" {
  rest_api_id = awscc_apigateway_rest_api.example.id
  stage_name  = "prod"

  tags = [{
    key   = "Modified By"
    value = "AWSCC"
  }]
}

# Create a domain name
resource "awscc_apigateway_domain_name" "example" {
  domain_name = "api.example.com"
  endpoint_configuration = {
    types = ["REGIONAL"]
  }
  regional_certificate_arn = aws_acm_certificate.example.arn

  tags = [{
    key   = "Modified By"
    value = "AWSCC"
  }]
}

# Create the base path mapping
resource "awscc_apigateway_base_path_mapping" "example" {
  domain_name = awscc_apigateway_domain_name.example.domain_name
  base_path   = "v1"
  rest_api_id = awscc_apigateway_rest_api.example.id
  stage       = awscc_apigateway_stage.example.stage_name
}