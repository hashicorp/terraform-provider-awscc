# Get current AWS region
data "aws_region" "current" {}

# Get current AWS account ID
data "aws_caller_identity" "current" {}

# Example MSK cluster policy to allow operations from trusted accounts
data "aws_iam_policy_document" "msk_policy" {
  statement {
    sid    = "AllowMSKOperations"
    effect = "Allow"
    principals {
      type        = "AWS"
      identifiers = [data.aws_caller_identity.current.account_id]
    }
    actions = [
      "kafka:DescribeCluster",
      "kafka:GetBootstrapBrokers"
    ]
    resources = ["*"]
  }
}

# Create MSK Cluster Policy
resource "awscc_msk_cluster_policy" "example" {
  cluster_arn = "arn:aws:kafka:${data.aws_region.current.name}:${data.aws_caller_identity.current.account_id}:cluster/example-cluster/12345678-abcd-1234-efgh-111122223333-2"
  policy      = jsonencode(jsondecode(data.aws_iam_policy_document.msk_policy.json))
}