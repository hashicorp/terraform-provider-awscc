resource "awscc_eks_addon" "coredns" {
  cluster_name = awscc_eks_cluster.eks_cluster.name
  addon_name   = "coredns"
  # Optional: addon_version = "v1.8.4-eksbuild.1"
  # Optional: resolve_conflicts = "OVERWRITE"
}

resource "awscc_eks_addon" "kube_proxy" {
  cluster_name = awscc_eks_cluster.eks_cluster.name
  addon_name   = "kube-proxy"
}
