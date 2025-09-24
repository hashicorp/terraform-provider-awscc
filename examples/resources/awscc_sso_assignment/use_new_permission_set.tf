data "aws_ssoadmin_instances" "example" {} // fetch IAM Identity Center instance arn

// create new permission set
resource "awscc_sso_permission_set" "example" {
  instance_arn = data.aws_ssoadmin_instances.example.arns[0] // reference existing IAM IDC instance by arn
  name         = "ExamplePermissionSet"                      // add desired name for permission set
  description  = "An example Permission Set"                 // add desired description for permission set
  // add multiple managed policies
  managed_policies = [
    "arn:aws:iam::aws:policy/job-function/ViewOnlyAccess",
  ]
  // redirect to S3 in us-east-1 upon sign-in
  relay_state_type = "https://s3.console.aws.amazon.com/s3/home?region=us-east-1#"
  // set 2 hour session duration
  session_duration = "PT2H"

  tags = [
    {
      key   = "Modified By"
      value = "AWSCC"
    }
  ]
}

data "aws_identitystore_group" "example" {
  identity_store_id = data.aws_ssoadmin_instances.example.identity_store_ids[0]

  alternate_identifier {
    unique_attribute {
      attribute_path  = "DisplayName"
      attribute_value = "ExampleGroup" // fetch info for existing IAM IDC group with DisplayName of 'ExampleGroup'
    }
  }
}

resource "awscc_sso_assignment" "example" {
  instance_arn       = data.aws_ssoadmin_instances.example.arns[0]
  permission_set_arn = awscc_sso_permission_set.example.permission_set_arn

  principal_id   = data.aws_identitystore_group.example.group_id // reference group id that was fetched by the data source
  principal_type = "GROUP"                                       // valid values are 'USER' or 'GROUP'

  target_id   = "012347678910"
  target_type = "AWS_ACCOUNT"
}
