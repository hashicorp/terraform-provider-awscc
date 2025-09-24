resource "awscc_eks_fargate_profile" "example" {
  cluster_name           = awscc_eks_cluster.example.name
  fargate_profile_name   = "example"
  pod_execution_role_arn = awscc_iam_role.example.arn
  subnets                = [awscc_ec2_subnet.example1.id, awscc_ec2_subnet.example2.id]

  selectors = [{
    namespace = "default"
    labels = [{
      key   = "env"
      value = "dev"
    }]
  }]
  tags = [
    {
      key   = "Managed By"
      value = "AWSCC"
    }
  ]
}