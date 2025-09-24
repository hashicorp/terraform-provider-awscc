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
