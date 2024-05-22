data "aws_caller_identity" "current" {}

resource "awscc_ecr_registry_policy" "example" {
  policy_text = jsonencode({
    Version = "2012-10-17"
    Statement = [
      {
        Sid    = "ReplicationAccessCrossAccount"
        Effect = "Allow"
        Principal = {
          AWS = "arn:aws:iam::${var.source_account}:root"
        }
        Action = [
          "ecr:CreateRepository",
          "ecr:ReplicateImage"
        ]
        Resource = "${awscc_ecr_repository.example.arn}/*"
      }
    ]
  })
}

resource "awscc_ecr_repository" "example" {
  repository_name      = "example-ecr"
  image_tag_mutability = "MUTABLE"
  image_scanning_configuration = {
    scan_on_push = true
  }
}

variable "source_account" {
  type = string
}