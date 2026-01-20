resource "awscc_iam_role" "observability" {
  role_name = "observability-admin-role"

  assume_role_policy_document = jsonencode({
    Version = "2012-10-17"
    Statement = [{
      Effect = "Allow"
      Principal = {
        Service = "observabilityadmin.amazonaws.com"
      }
      Action = "sts:AssumeRole"
    }]
  })

  tags = [{
    key   = "Environment"
    value = "test"
  }]
}

resource "awscc_observabilityadmin_s3_table_integration" "example" {
  role_arn = awscc_iam_role.observability.arn
  encryption = {
    sse_algorithm = "AES256"
  }

  tags = [{
    key   = "Environment"
    value = "test"
  }]
}
