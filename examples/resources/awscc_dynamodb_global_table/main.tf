data "aws_region" "current" {}
data "aws_region" "secondary" {
  provider = aws.secondary
}

# Example DynamoDB Global Table
resource "awscc_dynamodb_global_table" "example" {
  table_name = "example-global-table"

  billing_mode = "PAY_PER_REQUEST"

  attribute_definitions = [{
    attribute_name = "id"
    attribute_type = "S"
  }]

  key_schema = [{
    attribute_name = "id"
    key_type       = "HASH"
  }]

  # Enable server-side encryption
  sse_specification = {
    sse_enabled = true
    sse_type    = "KMS"
  }

  # Enable Time to Live
  time_to_live_specification = {
    enabled        = true
    attribute_name = "ttl"
  }

  # Enable DynamoDB Streams
  stream_specification = {
    stream_view_type = "NEW_AND_OLD_IMAGES"
  }

  # Define replicas
  replicas = [
    {
      region                      = data.aws_region.current.name
      deletion_protection_enabled = false
      point_in_time_recovery_specification = {
        point_in_time_recovery_enabled = true
      }
      tags = [{
        key   = "Modified By"
        value = "AWSCC"
      }]
    },
    {
      region                      = data.aws_region.secondary.name
      deletion_protection_enabled = false
      point_in_time_recovery_specification = {
        point_in_time_recovery_enabled = true
      }
      tags = [{
        key   = "Modified By"
        value = "AWSCC"
      }]
    }
  ]
}