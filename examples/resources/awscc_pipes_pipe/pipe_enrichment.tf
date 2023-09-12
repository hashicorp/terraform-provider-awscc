data "aws_iam_role" "example" {
  name = "awscc_example"
}

data "aws_cloudwatch_event_connection" "example" {
  name = "awscc-example"
}

resource "awscc_sqs_queue" "source" {}

resource "awscc_sqs_queue" "target" {}

resource "awscc_events_api_destination" "example" {
  connection_arn      = data.aws_cloudwatch_event_connection.example.arn
  http_method         = "GET"
  invocation_endpoint = "https://example.com"
}

resource "awscc_pipes_pipe" "example" {
  name       = "example-pipe"
  role_arn   = data.aws_iam_role.example.arn
  source     = awscc_sqs_queue.source.arn
  target     = awscc_sqs_queue.target.arn
  enrichment = awscc_events_api_destination.example.arn
  enrichment_parameters = {
    http_parameters = {
      path_parameter_values = ["example-path-param"]
      header_parameters = {
        "example-header"        = "example-value"
        "second-example-header" = "second-example-value"
      }
      query_string_parameteres = {
        "example-query-string"        = "example-value"
        "second-example-query-string" = "second-example-value"
      }
    }
  }
}