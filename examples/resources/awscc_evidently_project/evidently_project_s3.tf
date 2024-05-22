resource "awscc_evidently_project" "example" {
  name        = "Example"
  description = "Example Description"

  data_delivery = {
    s3_destination = {
      bucket_name = var.destination_bucket_name
      prefix      = "example"
    }
  }

  tags = [
    {
      key   = "ModifiedBy"
      value = "AWSCC"
    }
  ]
}

variable "destination_bucket_name" {
  type = string
}