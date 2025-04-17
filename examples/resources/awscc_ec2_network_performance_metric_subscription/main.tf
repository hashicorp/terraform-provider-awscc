resource "awscc_ec2_network_performance_metric_subscription" "example" {
  source      = "us-east-1" # Source region
  destination = "us-west-2" # Destination region (must be different from source)
  metric      = "aggregate-latency"
  statistic   = "p50"
}