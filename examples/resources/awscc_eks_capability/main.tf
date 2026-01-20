resource "awscc_eks_capability" "example" {
  cluster_name              = awscc_eks_cluster.example.name
  capability_name           = "example-capability"
  type                      = "ARGOCD"
  role_arn                  = awscc_iam_role.eks_capability_role.arn
  delete_propagation_policy = "RETAIN"

  tags = [{
    key   = "Environment"
    value = "test"
  }]
}
