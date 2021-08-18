terraform {
  required_providers {
    awscc = {
      source = "hashicorp/awscc"
    }
  }
}

provider "awscc" {
  region = "us-west-2"
}

resource "aws_kms_key" "test" {
  provider = awscc

  key_policy = jsonencode({
    Id = "kms-tf-1"
    Statement = [
      {
        Action = "kms:*"
        Effect = "Allow"
        Principal = {
          AWS = "*"
        }

        Resource = "*"
        Sid      = "Enable IAM User Permissions"
      },
    ]
    Version = "2012-10-17"
  })

  pending_window_in_days = 8

  # tags = [
  #   {
  #     key   = "Name"
  #     value = "Testing"
  #   },
  #   {
  #     key   = "Env"
  #     value = "dev"
  #   }
  # ]
}