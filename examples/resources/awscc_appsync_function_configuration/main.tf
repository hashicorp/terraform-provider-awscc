# Get current AWS region
data "aws_region" "current" {}

# Create AppSync API (no AWSCC equivalent yet)
resource "aws_appsync_graphql_api" "example" {
  name                = "example-api"
  authentication_type = "API_KEY"
  schema              = <<EOF
type Query {
  getExample(id: ID!): Example
}

type Example {
  id: ID!
  value: String
}
EOF

  tags = {
    "Modified By" = "AWSCC"
  }
}

# Example DynamoDB table for the data source
resource "awscc_dynamodb_table" "example" {
  attribute_definitions = [
    {
      attribute_name = "id"
      attribute_type = "S"
    }
  ]
  key_schema = [
    {
      attribute_name = "id"
      key_type       = "HASH"
    }
  ]
  table_name   = "example-table"
  billing_mode = "PAY_PER_REQUEST"
  tags = [{
    key   = "Modified By"
    value = "AWSCC"
  }]
}

# IAM role for AppSync to access DynamoDB
resource "awscc_iam_role" "appsync_datasource_role" {
  role_name = "example-appsync-dynamodb-role"
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

resource "awscc_iam_role_policy" "appsync_dynamodb_policy" {
  policy_name = "DynamoDBAccess"
  role_name   = awscc_iam_role.appsync_datasource_role.role_name
  policy_document = jsonencode({
    Version = "2012-10-17"
    Statement = [
      {
        Effect = "Allow"
        Action = [
          "dynamodb:GetItem",
          "dynamodb:PutItem",
          "dynamodb:DeleteItem",
          "dynamodb:UpdateItem",
          "dynamodb:Query",
          "dynamodb:Scan"
        ]
        Resource = [
          awscc_dynamodb_table.example.arn,
          "${awscc_dynamodb_table.example.arn}/index/*"
        ]
      }
    ]
  })
}

# Create AppSync datasource
resource "aws_appsync_datasource" "example" {
  api_id = aws_appsync_graphql_api.example.id
  name   = "example_datasource"
  type   = "AMAZON_DYNAMODB"

  dynamodb_config {
    table_name = awscc_dynamodb_table.example.table_name
    region     = data.aws_region.current.name
  }

  service_role_arn = awscc_iam_role.appsync_datasource_role.arn
}

# Create AppSync function configuration using AWSCC
resource "awscc_appsync_function_configuration" "example" {
  api_id           = aws_appsync_graphql_api.example.id
  data_source_name = aws_appsync_datasource.example.name
  name             = "exampleFunction"
  description      = "Example AppSync Function"

  runtime = {
    name            = "APPSYNC_JS"
    runtime_version = "1.0.0"
  }

  code = <<EOF
export function request(ctx) {
  return {
    operation: 'GetItem',
    key: util.dynamodb.toMapValues({id: ctx.args.id})
  };
}

export function response(ctx) {
  return ctx.result;
}
EOF

  function_version = "2018-05-29"
}