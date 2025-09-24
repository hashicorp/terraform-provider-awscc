resource "awscc_qbusiness_application" "example" {
  description                  = "Example QBusiness Application"
  display_name                 = "example_q_app"
  identity_center_instance_arn = data.aws_ssoadmin_instances.example.arns[0]
  attachments_configuration = {
    attachments_control_mode = "ENABLED"
  }

  tags = [{
    key   = "Modified By"
    value = "AWSCC"
  }]

}

data "aws_ssoadmin_instances" "example" {}