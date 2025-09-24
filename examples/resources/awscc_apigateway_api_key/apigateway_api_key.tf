resource "awscc_apigateway_api_key" "example" {
  name        = "example"
  description = "Example API key"
  enabled     = true

  tags = [{
    key   = "Modified By"
    value = "AWSCC"
  }]
}