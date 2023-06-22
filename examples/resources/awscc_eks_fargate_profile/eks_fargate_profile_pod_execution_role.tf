resource "awscc_iam_role" "example" {
  role_name   = "example-AmazonEKSFargatePodExecutionRole"
  description = "Example AWS FargatePod execution role"
  assume_role_policy_document = jsonencode({
    Statement = [{
      Action = "sts:AssumeRole"
      Effect = "Allow"
      Principal = {
        Service = "eks-fargate-pods.amazonaws.com"
      }
    }]
    Version = "2012-10-17"
  })
  managed_policy_arns = ["arn:aws:iam::aws:policy/AmazonEKSFargatePodExecutionRolePolicy"]
}