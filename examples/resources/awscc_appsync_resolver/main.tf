# Create the AppSync API
resource "awscc_appsync_graph_ql_api" "example" {
  name                = "example-api"
  authentication_type = "API_KEY"
  tags = [{
    key   = "Modified By"
    value = "AWSCC"
  }]
}

# Create AppSync Datasource
resource "awscc_appsync_data_source" "example" {
  api_id = awscc_appsync_graph_ql_api.example.id
  name   = "example_datasource"
  type   = "NONE"
}

# Create an AppSync Resolver
resource "awscc_appsync_resolver" "example" {
  api_id           = awscc_appsync_graph_ql_api.example.id
  type_name        = "Query"
  field_name       = "hello"
  data_source_name = awscc_appsync_data_source.example.name
  kind             = "UNIT"

  request_mapping_template = <<EOF
{
    "version": "2018-05-29",
    "payload": "Hello from AppSync!"
}
EOF

  response_mapping_template = "$util.toJson($context.result)"

  caching_config = {
    ttl          = 60
    caching_keys = ["$context.identity.sub"]
  }
}