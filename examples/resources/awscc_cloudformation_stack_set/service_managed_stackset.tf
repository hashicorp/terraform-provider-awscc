resource "awscc_cloudformation_stack_set" "stackset" {
  stack_set_name = "network-stackset1"

  permission_model = "SERVICE_MANAGED"

  auto_deployment = {
    enabled                          = true
    retain_stacks_on_account_removal = false
  }

  template_body = jsonencode({

    Resources = {
      myVpc = {
        Type = "AWS::EC2::VPC"
        Properties = {
          CidrBlock = "10.0.0.0/16"
          Tags = [
            {
              Key   = "Name"
              Value = "Primary_CF_VPC"
            }
          ]
        }
      }
    }
  })
}