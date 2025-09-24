data "aws_ssoadmin_instances" "example" {} // fetch IAM Identity Center instance arn

data "aws_ssoadmin_permission_set" "example" {
  instance_arn = data.aws_ssoadmin_instances.example.arns[0]
  name         = "AWSReadOnlyAccess" // fetch existing default 'AWSReadOnlyAccess' permission set
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
  permission_set_arn = data.aws_ssoadmin_permission_set.example.arn

  principal_id   = data.aws_identitystore_group.example.group_id // reference group id that was fetched by the data source
  principal_type = "GROUP"                                       // valid values are 'USER' or 'GROUP'

  target_id   = "012347678910"
  target_type = "AWS_ACCOUNT"


}
