# IAM role for AppSync merged API
resource "awscc_iam_role" "appsync_merged" {
  role_name = "appsync-merged-api-role"
  assume_role_policy_document = jsonencode({
    Version = "2012-10-17"
    Statement = [
      {
        Action = "sts:AssumeRole"
        Effect = "Allow"
        Principal = {
          Service = "appsync.amazonaws.com"
        }
      }
    ]
  })

  tags = [{
    key   = "Modified By"
    value = "AWSCC"
  }]
}

# IAM role policy for AppSync merged API
resource "awscc_iam_role_policy" "appsync_merged" {
  role_name   = awscc_iam_role.appsync_merged.role_name
  policy_name = "appsync-merged-api-policy"
  policy_document = jsonencode({
    Version = "2012-10-17"
    Statement = [
      {
        Effect = "Allow"
        Action = [
          "appsync:GetGraphqlApi",
          "appsync:GraphQL"
        ]
        Resource = [
          "*"
        ]
      }
    ]
  })
}

# Create source AppSync API
resource "aws_appsync_graphql_api" "source" {
  name                = "source-api"
  authentication_type = "API_KEY"
  schema              = "type Query { hello: String }"
  xray_enabled        = false
  api_type            = "GRAPHQL"

  tags = {
    "Modified By" = "AWSCC"
  }
}

# Create merged AppSync API
resource "aws_appsync_graphql_api" "merged" {
  name                          = "merged-api"
  authentication_type           = "API_KEY"
  xray_enabled                  = false
  api_type                      = "MERGED"
  merged_api_execution_role_arn = awscc_iam_role.appsync_merged.arn

  tags = {
    "Modified By" = "AWSCC"
  }
}

# Create the SourceApiAssociation
resource "awscc_appsync_source_api_association" "example" {
  source_api_identifier = aws_appsync_graphql_api.source.id
  merged_api_identifier = aws_appsync_graphql_api.merged.id
  description           = "Example source API association"
  source_api_association_config = {
    merge_type = "AUTO_MERGE"
  }
}