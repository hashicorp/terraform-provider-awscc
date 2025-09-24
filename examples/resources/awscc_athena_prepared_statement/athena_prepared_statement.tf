resource "awscc_athena_prepared_statement" "example" {
  query_statement = "SELECT * FROM ${var.athena_db_name} WHERE x = ?"
  statement_name  = "example"
  work_group      = awscc_athena_work_group.example.name
}

resource "awscc_athena_work_group" "example" {
  name        = "example"
  description = "Example Athena work group"

  work_group_configuration = {
    enforce_work_group_configuration   = true
    publish_cloudwatch_metrics_enabled = true

    result_configuration = {
      output_location = "s3://${var.s3_bucket_name}/output/"

      encryption_configuration = {
        encryption_option = "SSE_KMS"
        kms_key           = var.kms_key_arn
      }
    }
  }

  tags = [
    {
      key   = "ModifiedBy"
      value = "AWSCC"
    }
  ]

}

variable "athena_db_name" {
  type = string
}

variable "s3_bucket_name" {
  type = string
}

variable "kms_key_arn" {
  type = string
}





