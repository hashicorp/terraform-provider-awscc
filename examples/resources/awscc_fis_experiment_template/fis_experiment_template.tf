resource "awscc_fis_experiment_template" "example" {
  description = "example"
  role_arn    = var.fis_role_arn

  tags = {
    "Name" = "example"
  }

  stop_conditions = [{
    source = "none"
  }]

  actions = {
    action_1 = {
      name        = "example-action"
      action_id   = "aws:ec2:terminate-instances"
      description = "example action"

      targets = {
        "Instances" = "example-target"
      }
    }
  }


  targets = {
    example-target = {
      name           = "example-target"
      resource_type  = "aws:ec2:instance"
      selection_mode = "COUNT(1)"
      resource_tags = {
        "Name" = "fis"
      }
    }
  }

  experiment_options = {
    empty_target_resolution_mode = "fail"
  }

}

variable "fis_role_arn" {
  type        = string
  description = "Role ARN for FIS"
}