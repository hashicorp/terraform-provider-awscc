# Data sources for AWS account and region
data "aws_caller_identity" "current" {}
data "aws_region" "current" {}

# Create an EventBridge Schema Registry first
resource "awscc_eventschemas_registry" "example" {
  registry_name = "example-registry"

  tags = [{
    key   = "Modified By"
    value = "AWSCC"
  }]
}

# Define the registry policy using policy document data source
data "aws_iam_policy_document" "registry_policy" {
  statement {
    effect = "Allow"
    principals {
      type = "AWS"
      identifiers = [
        "arn:aws:iam::${data.aws_caller_identity.current.account_id}:root"
      ]
    }
    actions = [
      "schemas:CreateSchema",
      "schemas:DeleteSchema",
      "schemas:DescribeSchema",
      "schemas:UpdateSchema",
      "schemas:ListSchemas"
    ]
    resources = [
      "arn:aws:schemas:${data.aws_region.current.name}:${data.aws_caller_identity.current.account_id}:registry/${awscc_eventschemas_registry.example.registry_name}/*"
    ]
  }
}

# Create the registry policy
resource "awscc_eventschemas_registry_policy" "example" {
  registry_name = awscc_eventschemas_registry.example.registry_name
  policy        = jsonencode(jsondecode(data.aws_iam_policy_document.registry_policy.json))
}