resource "awscc_apigatewayv2_api" "example_http_api" {
  name          = "example-http-api"
  protocol_type = "HTTP"
  tags = {
    key   = "Modified By"
    value = "AWSCC"
  }
}