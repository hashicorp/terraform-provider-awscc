resource "awscc_iot_role_alias" "example" {
  role_arn                    = awscc_iam_role.example.arn
  role_alias                  = "example"
  credential_duration_seconds = 900

  tags = [
    {
      key   = "Modified By"
      value = "AWSCC"
    }
  ]
}

resource "awscc_iam_role" "example" {
  role_name   = "example"
  description = "Role that allows IoT to write to Cloudwatch logs"
  assume_role_policy_document = jsonencode({
    Version = "2012-10-17"
    Statement = [
      {
        Action = "sts:AssumeRole"
        Effect = "Allow"
        Principal = {
          Service = "iot.amazonaws.com"
        }
      }
    ]
  })
  tags = [
    {
      key   = "Modified By"
      value = "AWSCC"
    }
  ]
}
