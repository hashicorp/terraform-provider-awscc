resource "awscc_batch_scheduling_policy" "example" {
  name = "example-scheduling-policy"

  fairshare_policy = {
    compute_reservation = 1
    share_decay_seconds = 3600
  }

  tags = [{
    key   = "Environment"
    value = "dev"
  }, {
    key   = "Modified By"
    value = "AWSCC"
  }]
}

output "scheduling_policy_arn" {
  description = "ARN of the Batch Scheduling Policy"
  value       = awscc_batch_scheduling_policy.example.arn
}