resource "awscc_athena_named_query" "example" {
  name         = "example"
  work_group   = awscc_athena_work_group.example.id
  database     = var.athena_db_name
  query_string = "SELECT * FROM var.athena_db_name limit 10;"
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





