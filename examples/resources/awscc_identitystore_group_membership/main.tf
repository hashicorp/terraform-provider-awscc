# Get Identity Store ID
data "aws_ssoadmin_instances" "test" {}

# Create a Group
resource "awscc_identitystore_group" "test" {
  identity_store_id = tolist(data.aws_ssoadmin_instances.test.identity_store_ids)[0]
  display_name      = "TestGroup"
  description       = "Test Group for AWSCC example"
}

# Create a User in Identity Store
resource "aws_identitystore_user" "test" {
  identity_store_id = tolist(data.aws_ssoadmin_instances.test.identity_store_ids)[0]
  display_name      = "Test User"
  user_name         = "test.user"
  name {
    given_name  = "Test"
    family_name = "User"
  }
  emails {
    primary = true
    value   = "test.user@example.com"
  }
}

# Create Group Membership
resource "awscc_identitystore_group_membership" "test" {
  identity_store_id = tolist(data.aws_ssoadmin_instances.test.identity_store_ids)[0]
  group_id          = awscc_identitystore_group.test.id
  member_id = {
    user_id = aws_identitystore_user.test.user_id
  }
}