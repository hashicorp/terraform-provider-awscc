# Data source for the SSO Instance
data "aws_ssoadmin_instances" "example" {}

# The identity store group
resource "awscc_identitystore_group" "example" {
  identity_store_id = tolist(data.aws_ssoadmin_instances.example.identity_store_ids)[0]
  display_name      = "example-group"
  description       = "Example group created with AWSCC provider"
}