# Get current AWS region
data "aws_region" "current" {}

# IAM role policy document for Image Builder Lifecycle Policy
data "aws_iam_policy_document" "assume_role" {
  statement {
    actions = ["sts:AssumeRole"]
    effect  = "Allow"

    principals {
      type        = "Service"
      identifiers = ["imagebuilder.amazonaws.com"]
    }
  }
}

data "aws_iam_policy_document" "lifecycle_policy" {
  statement {
    effect = "Allow"
    actions = [
      "ec2:DeleteSnapshot",
      "ec2:DeregisterImage",
      "ec2:DescribeImages",
      "ec2:DescribeSnapshots",
      "imagebuilder:GetImage"
    ]
    resources = ["*"]
  }
}

# Create IAM role for Image Builder Lifecycle Policy
resource "awscc_iam_role" "lifecycle_policy" {
  assume_role_policy_document = data.aws_iam_policy_document.assume_role.json
  description                 = "Role for Image Builder Lifecycle Policy"
  path                        = "/"
  role_name                   = "ImageBuilderLifecyclePolicyRole"

  tags = [{
    key   = "Modified By"
    value = "AWSCC"
  }]
}

resource "awscc_iam_role_policy" "lifecycle_policy" {
  policy_document = data.aws_iam_policy_document.lifecycle_policy.json
  policy_name     = "ImageBuilderLifecyclePolicyPermissions"
  role_name       = awscc_iam_role.lifecycle_policy.role_name
}

# Create Image Builder Lifecycle Policy
resource "awscc_imagebuilder_lifecycle_policy" "example" {
  name           = "example-lifecycle-policy"
  description    = "Example Image Builder Lifecycle Policy"
  execution_role = awscc_iam_role.lifecycle_policy.arn
  resource_type  = "AMI_IMAGE"
  status         = "ENABLED"

  policy_details = [
    {
      action = {
        type = "DELETE"
        include_resources = {
          amis       = true
          containers = false
          snapshots  = true
        }
      }
      filter = {
        type  = "AGE"
        value = 30
        unit  = "DAYS"
      }
      exclusion_rules = {
        amis = {
          regions = [data.aws_region.current.name]
        }
        tag_map = {
          "Environment" = "Production"
        }
      }
    }
  ]

  resource_selection = {
    tag_map = {
      "Environment" = "Development"
    }
  }

  tags = [{
    key   = "Modified By"
    value = "AWSCC"
  }]
}