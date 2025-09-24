data "aws_caller_identity" "current" {}
data "aws_region" "current" {}

resource "awscc_iam_role" "dataset_role" {
  role_name = "iotsitewise-dataset-role"
  assume_role_policy_document = jsonencode({
    Version = "2012-10-17"
    Statement = [
      {
        Action = "sts:AssumeRole"
        Effect = "Allow"
        Principal = {
          Service = "iotsitewise.amazonaws.com"
        }
      }
    ]
  })

  tags = [{
    key   = "Modified By"
    value = "AWSCC"
  }]
}

resource "aws_iam_policy" "dataset_policy" {
  name = "iotsitewise-dataset-policy"
  policy = jsonencode({
    Version = "2012-10-17"
    Statement = [
      {
        Effect = "Allow"
        Action = [
          "kendra:ListTagsForResource",
          "kendra:GetKnowledgeBase",
          "kendra:DescribeKnowledgeBase"
        ]
        Resource = "*"
      }
    ]
  })

  tags = {
    "Modified By" = "AWSCC"
  }
}

resource "aws_iam_role_policy_attachment" "dataset_role_policy" {
  policy_arn = aws_iam_policy.dataset_policy.arn
  role       = awscc_iam_role.dataset_role.role_name
}

locals {
  knowledge_base_arn = "arn:aws:kendra:${data.aws_region.current.name}:${data.aws_caller_identity.current.account_id}:knowledgebase/example"
}

resource "awscc_iotsitewise_dataset" "example" {
  dataset_name        = "example-dataset"
  dataset_description = "Example IoT SiteWise Dataset"

  dataset_source = {
    source_type   = "KENDRA"
    source_format = "KNOWLEDGE_BASE"
    source_detail = {
      kendra = {
        knowledge_base_arn = local.knowledge_base_arn
        role_arn           = awscc_iam_role.dataset_role.arn
      }
    }
  }

  tags = [{
    key   = "Modified By"
    value = "AWSCC"
  }]
}