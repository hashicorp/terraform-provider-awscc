resource "awscc_cloudformation_stack_set" "stackset" {
  stack_set_name = "network-stackset"

  permission_model = "SELF_MANAGED"

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