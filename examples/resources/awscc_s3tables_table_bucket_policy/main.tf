data "aws_caller_identity" "current" {}

resource "awscc_s3tables_table_bucket" "example" {
  table_bucket_name = "example-table-bucket-${data.aws_caller_identity.current.account_id}"
}

data "aws_iam_policy_document" "table_bucket_policy" {
  statement {
    sid    = "AllowTableAccess"
    effect = "Allow"

    principals {
      type        = "AWS"
      identifiers = ["arn:aws:iam::${data.aws_caller_identity.current.account_id}:root"]
    }

    actions = [
      "s3tables:GetTable",
      "s3tables:ListTables"
    ]

    resources = ["${awscc_s3tables_table_bucket.example.table_bucket_arn}/*"]
  }
}

# S3 Tables table bucket policy
resource "awscc_s3tables_table_bucket_policy" "example" {
  table_bucket_arn = awscc_s3tables_table_bucket.example.table_bucket_arn
  resource_policy  = data.aws_iam_policy_document.table_bucket_policy.json
}
