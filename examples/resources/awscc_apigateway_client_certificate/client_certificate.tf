resource "awscc_apigateway_client_certificate" "example" {
  description = "My API Gateway client certificate"
  tags = [{
    key   = "Managed By"
    value = "AWSCC"
  }]
}
