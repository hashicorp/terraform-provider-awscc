resource "awscc_kendra_faq" "example" {
  index_id = var.kendra_index_id
  name     = "Example"
  role_arn = awscc_iam_role.faq.arn

  s3_path = {
    bucket = var.bucket_name
    key    = var.faq_key
  }

  description   = "Example of Kendra FAQ"
  file_format   = "CSV"
  language_code = "en"

  tags = [{
    key   = "Modified By"
    value = "AWSCC"
  }]
}

resource "awscc_iam_role" "faq" {
  role_name   = "kendra_faq_role"
  description = "Role used for Kendra FAQs"
  assume_role_policy_document = jsonencode({
    Version = "2012-10-17"
    Statement = [
      {
        Action = "sts:AssumeRole"
        Effect = "Allow"
        Principal = {
          Service = "kendra.amazonaws.com"
        }
      }
    ]
  })
  tags = [
    {
      key   = "Name"
      value = "Kendra FAQ role"
    },
    {
      key   = "Modified By"
      value = "AWSCC"
    }
  ]
}

resource "awscc_iam_role_policy" "faq_policy" {
  policy_name = "kendra_faq_role_policy"
  role_name   = awscc_iam_role.faq.id

  policy_document = jsonencode({
    Version = "2012-10-17"
    Statement = [
      {
        Effect   = "Allow"
        Action   = "s3:GetObject"
        Resource = "arn:aws:s3:::${var.bucket_name}/*"
      },

      {
        Effect   = "Allow"
        Action   = "kms:Decrypt",
        Resource = "arn:aws:kms:your-region:your-account-id:key/key-id",
        Condition = {
          "StringLike" : {
            "kms:ViaService" : "kendra.your-region.amazonaws.com"
          }
        }
      },

    ]
  })
}

variable "kendra_index_id" {
  type        = string
  description = "Kendra index id"
}

variable "bucket_name" {
  type        = string
  description = "S3 Bucket holding the FAQ document"
}

variable "faq_key" {
  type        = string
  description = "Bucket key for the FAQ document"
}
