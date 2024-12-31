resource "awscc_eks_addon" "coredns" {
  cluster_name = var.cluster_name
  addon_name   = "coredns"
  # Optional: addon_version = "v1.8.4-eksbuild.1"
  # Optional: resolve_conflicts = "OVERWRITE"
}

resource "awscc_eks_addon" "kube_proxy" {
  cluster_name = var.cluster_name
  addon_name   = "kube-proxy"
}

variable "cluster_name" {
  type = string
}
