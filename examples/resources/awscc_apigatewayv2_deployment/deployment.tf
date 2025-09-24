resource "awscc_apigatewayv2_deployment" "example" {
  api_id      = awscc_apigatewayv2_api.example.id
  description = "Beta deployment"
  stage_name  = "beta"
}
