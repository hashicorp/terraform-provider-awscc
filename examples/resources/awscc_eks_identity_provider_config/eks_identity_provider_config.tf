resource "awscc_eks_identity_provider_config" "identity_provider_config" {
  cluster_name                  = awscc_eks_cluster.eks_cluster.name
  type                          = "oidc"
  identity_provider_config_name = "identity_provider_config"

  oidc {
    client_id  = "sts.amazonaws.com"
    issuer_url = "https://${awscc_iam_oidc_provider.example.url}"
  }

  tags = [{
    key   = "Modified By"
    value = "AWSCC"
  }]
}
