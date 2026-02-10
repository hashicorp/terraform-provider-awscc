resource "awscc_quicksight_action_connector" "example" {
  action_connector_name = "example-action-connector"
  aws_account_id        = "012345678901"

  action_connector_definition = {
    name        = "example-action-connector"
    type        = "GENERIC_HTTP"
    description = "Example Generic HTTP action connector for QuickSight"

    authentication_config = {
      secret_arn          = "arn:aws:secretsmanager:us-west-2:012345678901:secret:quicksight-connector-secret"
      authentication_type = "SECRETSMANAGER"
    }

    url_template = "https://example.com/webhook"
  }
}

