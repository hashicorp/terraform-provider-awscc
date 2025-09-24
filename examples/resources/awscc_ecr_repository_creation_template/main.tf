data "aws_caller_identity" "current" {}

data "aws_iam_policy_document" "ecr_template_policy" {
  statement {
    sid    = "AllowPull"
    effect = "Allow"

    principals {
      type = "AWS"
      identifiers = [
        "arn:aws:iam::${data.aws_caller_identity.current.account_id}:root"
      ]
    }

    actions = [
      "ecr:GetDownloadUrlForLayer",
      "ecr:BatchGetImage",
      "ecr:BatchCheckLayerAvailability"
    ]
  }
}

# Create an IAM role for ECR repository template
resource "awscc_iam_role" "ecr_template_role" {
  role_name = "ECRTemplateRole"
  assume_role_policy_document = jsonencode({
    Version = "2012-10-17"
    Statement = [
      {
        Action = "sts:AssumeRole"
        Effect = "Allow"
        Principal = {
          Service = "ecr.amazonaws.com"
        }
      }
    ]
  })

  tags = [{
    key   = "Modified By"
    value = "AWSCC"
  }]
}

# Create the ECR repository creation template
resource "awscc_ecr_repository_creation_template" "example" {
  prefix      = "demo"
  description = "Demo template for ECR repositories"
  applied_for = ["REPLICATION"]

  custom_role_arn      = awscc_iam_role.ecr_template_role.arn
  repository_policy    = jsonencode(jsondecode(data.aws_iam_policy_document.ecr_template_policy.json))
  image_tag_mutability = "IMMUTABLE"
  lifecycle_policy = jsonencode({
    rules = [
      {
        rulePriority = 1
        description  = "Keep only 10 images"
        selection = {
          tagStatus   = "untagged"
          countType   = "imageCountMoreThan"
          countNumber = 10
        }
        action = {
          type = "expire"
        }
      }
    ]
  })

  encryption_configuration = {
    encryption_type = "AES256"
  }

  resource_tags = [{
    key   = "Modified By"
    value = "AWSCC"
  }]
}