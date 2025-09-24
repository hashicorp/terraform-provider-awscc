resource "awscc_evidently_project" "example" {
  name        = "Example"
  description = "Example Description"

  data_delivery = {
    log_group = var.log_group_name
  }

  tags = [
    {
      key   = "ModifiedBy"
      value = "AWSCC"
    }
  ]
}

variable "log_group_name" {
  type = string
}