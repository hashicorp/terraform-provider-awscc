# Data sources to get AWS account ID
data "aws_caller_identity" "current" {}

resource "awscc_s3tables_table_bucket" "example" {
  table_bucket_name = "my-s3tables-table-bucket-${data.aws_caller_identity.current.account_id}"

  unreferenced_file_removal = {
    status            = "Enabled"
    unreferenced_days = 7
    noncurrent_days   = 30
  }
}