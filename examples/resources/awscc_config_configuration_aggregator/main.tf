# Use data source for region
data "aws_region" "current" {}

# IAM role for Config Aggregator
data "aws_iam_policy_document" "assume_role" {
  statement {
    effect = "Allow"
    principals {
      type        = "Service"
      identifiers = ["config.amazonaws.com"]
    }
    actions = ["sts:AssumeRole"]
  }
}

data "aws_iam_policy_document" "config_aggregator_policy" {
  statement {
    effect = "Allow"
    actions = [
      "organizations:ListAccounts",
      "organizations:DescribeOrganization",
      "organizations:ListAWSServiceAccessForOrganization"
    ]
    resources = ["*"]
  }
}

resource "awscc_iam_role" "config_aggregator_role" {
  role_name                   = "AWSConfigAggregatorRole"
  assume_role_policy_document = data.aws_iam_policy_document.assume_role.json
  policies = [{
    policy_document = data.aws_iam_policy_document.config_aggregator_policy.json
    policy_name     = "ConfigAggregatorPolicy"
  }]
  tags = [{
    key   = "Modified By"
    value = "AWSCC"
  }]
}

# AWS Config Configuration Aggregator
resource "awscc_config_configuration_aggregator" "example" {
  configuration_aggregator_name = "example-aggregator"

  organization_aggregation_source = {
    role_arn        = awscc_iam_role.config_aggregator_role.arn
    aws_regions     = [data.aws_region.current.name]
    all_aws_regions = false
  }

  tags = [{
    key   = "Modified By"
    value = "AWSCC"
  }]
}