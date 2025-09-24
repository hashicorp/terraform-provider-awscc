resource "awscc_route53recoverycontrol_cluster" "example" {
  name = "example-recovery-control-cluster"

  tags = [{
    key   = "Modified By"
    value = "AWSCC"
  }]
}

output "cluster_arn" {
  value = awscc_route53recoverycontrol_cluster.example.cluster_arn
}

output "cluster_endpoints" {
  value = awscc_route53recoverycontrol_cluster.example.cluster_endpoints
}