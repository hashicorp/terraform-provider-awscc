resource "aws_iam_role" "example" {
  name = "example-servicecatalog-role"

  assume_role_policy = jsonencode({
    Version = "2012-10-17"
    Statement = [
      {
        Action = "sts:AssumeRole"
        Effect = "Allow"
        Principal = {
          Service = "servicecatalog.amazonaws.com"
        }
      }
    ]
  })

  tags = {
    Name        = "example-servicecatalog-role"
    Environment = "example"
  }
}

resource "aws_servicecatalog_portfolio" "example" {
  name         = "example-portfolio"
  description  = "Example Service Catalog portfolio"
  provider_name = "example-provider"

  tags = {
    Name        = "example-portfolio"
    Environment = "example"
  }
}

resource "awscc_servicecatalog_portfolio_principal_association" "example" {
  portfolio_id   = aws_servicecatalog_portfolio.example.id
  principal_arn  = aws_iam_role.example.arn
  principal_type = "IAM"
}
