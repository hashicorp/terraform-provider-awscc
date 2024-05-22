resource "awscc_athena_capacity_reservation" "example" {
  name        = "example"
  target_dpus = 24

  capacity_assignment_configuration = {
    capacity_assignments = [{
      workgroup_names = [var.athena_workgroup_name]
    }]
  }

  tags = [
    {
      key   = "ModifiedBy"
      value = "AWSCC"
    }
  ]

}

variable "athena_workgroup_name" {
  type = string
}
