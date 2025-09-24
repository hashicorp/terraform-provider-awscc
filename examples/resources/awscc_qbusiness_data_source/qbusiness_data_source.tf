resource "awscc_qbusiness_data_source" "example" {
  application_id = awscc_qbusiness_application.example.application_id
  display_name   = "example_q_data_source"
  index_id       = awscc_qbusiness_index.example.index_id
  role_arn       = awscc_iam_role.example.arn
  configuration = jsonencode(
    {
      type     = "S3"
      version  = "1.0.0"
      syncMode = "FORCED_FULL_CRAWL"
      connectionConfiguration = {
        repositoryEndpointMetadata = {
          BucketName = var.bucket_name
        }
      }
      additionalProperties = {
        inclusionPrefixes = ["docs/"]
      }
      repositoryConfigurations = {
        document = {
          fieldMappings = [
            {
              dataSourceFieldName = "s3_document_id"
              indexFieldType      = "STRING"
              indexFieldName      = "s3_document_id"
            }
          ]
        }
      }
    }
  )
  tags = [{
    key   = "Modified By"
    value = "AWSCC"
  }]
}

resource "awscc_iam_role" "example" {
  role_name   = "QBusiness-DataSource-Role"
  description = "QBusiness Data source role"
  assume_role_policy_document = jsonencode({
    Version = "2012-10-17"
    Statement = [
      {
        Sid    = "AllowsAmazonQToAssumeRoleForServicePrincipal"
        Effect = "Allow"
        Principal = {
          Service = "qbusiness.amazonaws.com"
        }
        Action = [
          "sts:AssumeRole"
        ]
        Condition = {
          StringEquals = {
            "aws:SourceAccount" = data.aws_caller_identity.current.account_id
          }
        }
      }
    ]
  })

  tags = [{
    key   = "Modified By"
    value = "AWSCC"
  }]
}

resource "awscc_iam_role_policy" "example" {
  policy_name = "sample_iam_role_policy"
  role_name   = awscc_iam_role.example.id

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
        Action   = "s3:ListBucket"
        Resource = "arn:aws:s3:::${var.bucket_name}"
      },
      {
        Effect = "Allow"
        Action = [
          "qbusiness:BatchPutDocument",
          "qbusiness:BatchDeleteDocument"
        ]
        Resource = "arn:aws:qbusiness:${data.aws_region.current.name}:${data.aws_caller_identity.current.account_id}:application/${awscc_qbusiness_application.example.id}/index/${awscc_qbusiness_index.example.id}"
      },
      {
        Effect = "Allow"
        Action = ["qbusiness:PutGroup",
          "qbusiness:CreateUser",
          "qbusiness:DeleteGroup",
          "qbusiness:UpdateUser",
        "qbusiness:ListGroups"]
        Resource = [
          "arn:aws:qbusiness:${data.aws_region.current.name}:${data.aws_caller_identity.current.account_id}:application/${awscc_qbusiness_application.example.id}",
          "arn:aws:qbusiness:${data.aws_region.current.name}:${data.aws_caller_identity.current.account_id}:application/${awscc_qbusiness_application.example.id}/index/${awscc_qbusiness_index.example.id}",
          "arn:aws:qbusiness:${data.aws_region.current.name}:${data.aws_caller_identity.current.account_id}:application/${awscc_qbusiness_application.example.id}/index/${awscc_qbusiness_index.example.id}/data-source/*"
        ]
      }
    ]
  })
}

variable "bucket_name" {
  type        = string
  description = "Name of the bucket to be used as the data source input"
}

data "aws_caller_identity" "current" {}

data "aws_region" "current" {}
