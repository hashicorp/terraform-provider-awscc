resource "awscc_eks_addon" "coredns" {
  cluster_name      = var.cluster_name
  addon_name        = "coredns"
  addon_version     = "v1.11.3-eksbuild.1"
  resolve_conflicts = "OVERWRITE"
  tags = [{
    key   = "Modified By"
    value = "AWSCC"
  }]
}

resource "awscc_eks_addon" "kube_proxy" {
  cluster_name = var.cluster_name
  addon_name   = "kube-proxy"
  # Optional: addon_version = "v1.30.5-eksbuild.2"
  # Optional: resolve_conflicts = "OVERWRITE"
  tags = [{
    key   = "Modified By"
    value = "AWSCC"
  }]
}


variable "cluster_name" {
  type = string
}
