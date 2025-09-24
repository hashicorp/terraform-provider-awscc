resource "awscc_apigatewayv2_vpc_link" "example" {
  name               = "example"
  subnet_ids         = [var.subnet_id]
  security_group_ids = [var.securit_group_id]
  tags = [{
    key   = "Managed By"
    value = "AWSCC"
  }]
}
