
# Get current AWS region
data "aws_region" "current" {}

# Get current AWS account ID
data "aws_caller_identity" "current" {}

# Create DynamoDB table with pay-per-request billing
resource "awscc_dynamodb_table" "example" {
  table_name   = "awscc-example-table"
  billing_mode = "PAY_PER_REQUEST"

  # Define primary key schema
  attribute_definitions = [
    {
      attribute_name = "UserId"
      attribute_type = "S"
    },
    {
      attribute_name = "GameTitle"
      attribute_type = "S"
    },
    {
      attribute_name = "TopScore"
      attribute_type = "N"
    }
  ]

  key_schema = jsonencode([
    {
      AttributeName = "UserId"
      KeyType       = "HASH"
    },
    {
      AttributeName = "GameTitle"
      KeyType       = "RANGE"
    }
  ])

  # Enable global secondary index
  global_secondary_indexes = [
    {
      index_name = "GameTitleIndex"
      key_schema = [
        {
          attribute_name = "GameTitle"
          key_type       = "HASH"
        },
        {
          attribute_name = "TopScore"
          key_type       = "RANGE"
        }
      ]
      projection = {
        projection_type = "ALL"
      }
    }
  ]

  # Enable point-in-time recovery
  point_in_time_recovery_specification = {
    point_in_time_recovery_enabled = true
  }

  # Enable server-side encryption with AWS managed key
  sse_specification = {
    sse_enabled = true
  }

  # Enable DynamoDB Streams
  stream_specification = {
    stream_view_type = "NEW_AND_OLD_IMAGES"
  }

  # Add table policy
  resource_policy = {
    policy_document = jsonencode({
      Version = "2012-10-17"
      Statement = [
        {
          Effect = "Allow"
          Principal = {
            AWS = "arn:aws:iam::${data.aws_caller_identity.current.account_id}:root"
          }
          Action = [
            "dynamodb:PutItem",
            "dynamodb:GetItem",
            "dynamodb:DeleteItem"
          ]
          Resource = "arn:aws:dynamodb:${data.aws_region.current.name}:${data.aws_caller_identity.current.account_id}:table/awscc-example-table"
        }
      ]
    })
  }

  tags = [
    {
      key   = "Environment"
      value = "Dev"
    },
    {
      key   = "Modified By"
      value = "AWSCC"
    }
  ]

  deletion_protection_enabled = false
}