resource "aws_iam_role" "example_glue_integration_role" {
  name = "example-glue-integration-role"

  assume_role_policy = jsonencode({
    Version = "2012-10-17"
    Statement = [
      {
        Action = "sts:AssumeRole"
        Effect = "Allow"
        Principal = {
          Service = "glue.amazonaws.com"
        }
      }
    ]
  })

  managed_policy_arns = [
    "arn:aws:iam::aws:policy/service-role/AWSGlueServiceRole"
  ]

  tags = {
    Name        = "example-glue-integration-role"
    Environment = "example"
  }
}

resource "aws_glue_catalog_database" "example_database" {
  name        = "example-integration-database"
  description = "Example database for Glue integration"

  tags = {
    Name        = "example-integration-database"
    Environment = "example"
  }
}

resource "aws_glue_connection" "example_connection" {
  name = "example-integration-connection"

  connection_properties = {
    JDBC_CONNECTION_URL = "jdbc:mysql://example-host:3306/example_db"
    USERNAME            = "example_user"
    PASSWORD            = "example_password"
  }

  connection_type = "JDBC"

  tags = {
    Name        = "example-integration-connection"
    Environment = "example"
  }
}

resource "awscc_glue_integration_resource_property" "example" {
  resource_arn = aws_glue_catalog_database.example_database.arn

  target_processing_properties = {
    connection_name = aws_glue_connection.example_connection.name
    role_arn        = aws_iam_role.example_glue_integration_role.arn
  }

  tags = [
    {
      key   = "Name"
      value = "example-integration-resource-property"
    },
    {
      key   = "Environment"
      value = "example"
    }
  ]
}