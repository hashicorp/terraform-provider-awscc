# Get current AWS region
data "aws_region" "current" {}

# Get current AWS account ID
data "aws_caller_identity" "current" {}

# Create AppSync API first
resource "aws_appsync_graphql_api" "example" {
  name                = "example-api"
  authentication_type = "API_KEY"

  tags = {
    "Modified By" = "AWSCC"
  }
}

# Create DynamoDB table that will be used as data source
resource "aws_dynamodb_table" "example" {
  name         = "example-table"
  billing_mode = "PAY_PER_REQUEST"
  hash_key     = "id"

  attribute {
    name = "id"
    type = "S"
  }

  tags = {
    "Modified By" = "AWSCC"
  }
}

# Create IAM role for AppSync to access DynamoDB
data "aws_iam_policy_document" "assume_role" {
  statement {
    effect = "Allow"
    principals {
      type        = "Service"
      identifiers = ["appsync.amazonaws.com"]
    }
    actions = ["sts:AssumeRole"]
  }
}

data "aws_iam_policy_document" "dynamodb_access" {
  statement {
    effect = "Allow"
    actions = [
      "dynamodb:GetItem",
      "dynamodb:PutItem",
      "dynamodb:DeleteItem",
      "dynamodb:UpdateItem",
      "dynamodb:Query",
      "dynamodb:Scan",
      "dynamodb:BatchGetItem",
      "dynamodb:BatchWriteItem"
    ]
    resources = [aws_dynamodb_table.example.arn]
  }
}

resource "aws_iam_role" "appsync_dynamodb_role" {
  name               = "AppSyncDynamoDBRole-Example"
  description        = "Role for AppSync to access DynamoDB"
  assume_role_policy = data.aws_iam_policy_document.assume_role.json

  inline_policy {
    name   = "AppSyncDynamoDBAccess"
    policy = data.aws_iam_policy_document.dynamodb_access.json
  }

  tags = {
    "Modified By" = "AWSCC"
  }
}

# Create AppSync data source
resource "aws_appsync_datasource" "example" {
  api_id = aws_appsync_graphql_api.example.id
  name   = "example_dynamodb_source"
  type   = "AMAZON_DYNAMODB"

  description = "Example DynamoDB data source"

  dynamodb_config {
    table_name = aws_dynamodb_table.example.name
    region     = data.aws_region.current.name
  }

  service_role_arn = aws_iam_role.appsync_dynamodb_role.arn
}