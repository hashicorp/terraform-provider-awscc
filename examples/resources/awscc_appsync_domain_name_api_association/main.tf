data "aws_caller_identity" "current" {}
data "aws_region" "current" {}

# Create an AppSync API
resource "awscc_appsync_graph_ql_api" "example" {
  name                = "example-api"
  authentication_type = "API_KEY"
  tags = [{
    key   = "Modified By"
    value = "AWSCC"
  }]
}

# Create a domain name for AppSync
resource "awscc_appsync_domain_name" "example" {
  domain_name     = "api.example.com"
  certificate_arn = "arn:aws:acm:${data.aws_region.current.name}:${data.aws_caller_identity.current.account_id}:certificate/example-certificate"
}

# Create the domain name API association
resource "awscc_appsync_domain_name_api_association" "example" {
  api_id      = awscc_appsync_graph_ql_api.example.id
  domain_name = awscc_appsync_domain_name.example.domain_name
}