data "aws_region" "current" {}
data "aws_caller_identity" "current" {}

data "aws_iam_policy_document" "assume_role" {
  statement {
    effect = "Allow"
    principals {
      type        = "Service"
      identifiers = ["comprehend.amazonaws.com"]
    }
    actions = ["sts:AssumeRole"]
  }
}

data "aws_iam_policy_document" "comprehend_flywheel" {
  statement {
    effect = "Allow"
    actions = [
      "s3:GetObject",
      "s3:ListBucket",
      "s3:PutObject"
    ]
    resources = [
      "arn:aws:s3:::example-datalake-${data.aws_caller_identity.current.account_id}-${data.aws_region.current.name}",
      "arn:aws:s3:::example-datalake-${data.aws_caller_identity.current.account_id}-${data.aws_region.current.name}/*"
    ]
  }

  statement {
    effect = "Allow"
    actions = [
      "comprehend:DescribeFlywheelIteration",
      "comprehend:ListFlywheelIterations"
    ]
    resources = ["*"]
  }
}

resource "awscc_iam_role" "comprehend_flywheel" {
  role_name                   = "comprehend-flywheel-role"
  assume_role_policy_document = data.aws_iam_policy_document.assume_role.json

  policies = [{
    policy_document = data.aws_iam_policy_document.comprehend_flywheel.json
    policy_name     = "comprehend-flywheel-policy"
  }]

  tags = [{
    key   = "Modified By"
    value = "AWSCC"
  }]
}

resource "awscc_comprehend_flywheel" "example" {
  flywheel_name        = "example-flywheel"
  data_lake_s3_uri     = "s3://example-datalake-${data.aws_caller_identity.current.account_id}-${data.aws_region.current.name}/flywheel-input"
  data_access_role_arn = awscc_iam_role.comprehend_flywheel.arn
  model_type           = "DOCUMENT_CLASSIFIER"

  task_config = {
    language_code = "en"
    document_classification_config = {
      mode   = "MULTI_LABEL"
      labels = ["POSITIVE", "NEGATIVE", "NEUTRAL"]
    }
  }

  tags = [{
    key   = "Modified By"
    value = "AWSCC"
  }]
}