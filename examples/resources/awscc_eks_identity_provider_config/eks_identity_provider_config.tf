resource "awscc_eks_identity_provider_config" "identity_provider_config" {
  cluster_name = awscc_eks_cluster.eks_cluster.name
  type         = "oidc"
  identity_provider_config_name = "identity_provider_config"

  oidc {
    client_id = "sts.amazonaws.com"
    issuer_url = aws_eks_cluster.eks_cluster.identity[0].oidc[0].issuer

  } # oid   
  tags = [{
    key = "Modified By"
    value = "AWSCC"
  }]
  # depends_on = [aws_eks_cluster.eks_cluster]

} 
