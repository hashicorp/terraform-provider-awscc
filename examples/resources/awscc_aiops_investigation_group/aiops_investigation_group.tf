# Create an IAM role for AIOps
resource "aws_iam_role" "aiops_role" {
  name = "aiops-investigation-role"

  assume_role_policy = jsonencode({
    Version = "2012-10-17"
    Statement = [
      {
        Action = "sts:AssumeRole"
        Effect = "Allow"
        Principal = {
          Service = "aiops.amazonaws.com"
        }
      }
    ]
  })

  # Add inline policy 
  inline_policy {
    name = "aiops-policy"
    policy = jsonencode({
      Version = "2012-10-17"
      Statement = [
        {
          Action = [
            "logs:*",
            "cloudtrail:*"
          ]
          Effect   = "Allow"
          Resource = "*"
        }
      ]
    })
  }
}

# Wait for IAM role to be fully created and propagated
resource "time_sleep" "wait_30_seconds" {
  depends_on      = [aws_iam_role.aiops_role]
  create_duration = "30s"
}

# AIOps Investigation Group resource with dependency on role
resource "awscc_aiops_investigation_group" "example" {
  depends_on = [time_sleep.wait_30_seconds]

  name                                 = "example-investigation-group"
  retention_in_days                    = 30
  is_cloud_trail_event_history_enabled = true
  role_arn                             = aws_iam_role.aiops_role.arn

  tags = [
    {
      key   = "Environment"
      value = "Development"
    },
    {
      key   = "Name"
      value = "example-investigation-group"
    }
  ]
}
