resource "awscc_securityhub_connector_v2" "example" {
  name          = "example-jira-connector"
  description   = "Example Security Hub Connector for Jira Cloud integration"
  provider_type = "JiraCloud"

  provider_configuration = {
    jira_cloud = {
      hostname    = "example.atlassian.net"
      project_key = "SEC"
      secret_arn  = "arn:aws:secretsmanager:us-west-2:123456789012:secret:jira-credentials-abcdef"
    }
  }
}
