# Configure Auth0 as an OIDC identity provider for EKS user authentication
# This allows users to authenticate to the EKS cluster using Auth0 credentials
resource "awscc_eks_identity_provider_config" "auth0_idp" {
  cluster_name = awscc_eks_cluster.eks_cluster.name
  type         = "oidc"

  oidc = {
    client_id      = var.oicd_client_id
    issuer_url     = var.oicd_issuer_url # Like: "https://dev-xxxxxxxxx.au.auth0.com"
    groups_claim   = "groups"
    username_claim = "email"
    groups_prefix  = var.oicd_groups_prefix # Like: "auth0:eks-cluster" 
  }

  tags = [
    {
      key   = "Provider"
      value = "Auth0"
    },
    {
      key   = "Modified By"
      value = "AWSCC"
    }
  ]

}
