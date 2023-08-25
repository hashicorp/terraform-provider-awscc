
resource "awscc_chatbot_slack_channel_configuration" "example" {
  # Set the name to identify the configuration in the ChatBot console
  configuration_name = "ChatBot-Example"

  # Replace <account_id> and <role-name> with sutable values
  iam_role_arn       = "arn:aws:iam::<account_id>:role/<role-name>"

  # Replace <channel-id> with the id of the Slack channel you want to send message to
  slack_channel_id   = "<channel-id>"

  # Replease <workspace-id> with the id of the Slack workspace
  slack_workspace_id = "<workspace-id>"
}
