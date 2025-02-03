# Get current AWS region
data "aws_region" "current" {}

# Get current AWS account ID
data "aws_caller_identity" "current" {}

# Model Package Group policy data source
data "aws_iam_policy_document" "model_package_group_policy" {
  statement {
    sid = "ModelPackageAccess"
    actions = [
      "sagemaker:DescribeModelPackage",
      "sagemaker:ListModelPackages",
      "sagemaker:CreateModelPackage",
    ]
    resources = [
      "arn:aws:sagemaker:${data.aws_region.current.name}:${data.aws_caller_identity.current.account_id}:model-package/${local.model_package_group_name}/*"
    ]
    principals {
      type        = "AWS"
      identifiers = [data.aws_caller_identity.current.account_id]
    }
  }
}

locals {
  model_package_group_name = "example-mpg"
}

resource "awscc_sagemaker_model_package_group" "example" {
  model_package_group_name        = local.model_package_group_name
  model_package_group_description = "Example SageMaker Model Package Group"
  model_package_group_policy      = jsonencode(jsondecode(data.aws_iam_policy_document.model_package_group_policy.json))

  tags = [{
    key   = "Modified By"
    value = "AWSCC"
  }]
}