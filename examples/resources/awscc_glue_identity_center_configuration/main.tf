# Get the SSO instance
data "aws_ssoadmin_instances" "example" {}

resource "awscc_glue_identity_center_configuration" "example" {
  instance_arn = tolist(data.aws_ssoadmin_instances.example.arns)[0]
}
