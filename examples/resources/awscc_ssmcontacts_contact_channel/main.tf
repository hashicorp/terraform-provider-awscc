data "aws_caller_identity" "current" {}
data "aws_region" "current" {}

# Create a contact channel
resource "awscc_ssmcontacts_contact_channel" "example" {
  channel_name    = "example-notification"
  channel_type    = "EMAIL"
  channel_address = "example@example.com"
  # In a real environment, you would need to provide an existing contact ARN
  contact_id       = "arn:aws:ssm-contacts:${data.aws_region.current.name}:${data.aws_caller_identity.current.account_id}:contact/example-contact"
  defer_activation = false
}