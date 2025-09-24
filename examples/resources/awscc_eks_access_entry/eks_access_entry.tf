resource "awscc_eks_access_entry" "example" {
  cluster_name  = var.cluster_name
  principal_arn = var.role_arn
  type          = "STANDARD"
  access_policies = [
    {
      access_scope = {
        type       = "namespace"
        namespaces = ["default"]
      }
      policy_arn = "arn:aws:eks::aws:cluster-access-policy/AmazonEKSAdminPolicy"
    },
    {
      access_scope = {
        type = "cluster"
      },
      policy_arn = "arn:aws:eks::aws:cluster-access-policy/AmazonEKSEditPolicy"
    }
  ]

  tags = [{
    key   = "Modified By"
    value = "AWSCC"
  }]
}

variable "cluster_name" {
  type = string
}

variable "role_arn" {
  type = string
}