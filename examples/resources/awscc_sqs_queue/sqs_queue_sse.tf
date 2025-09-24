resource "awscc_sqs_queue" "terraform_awscc_queue_sse" {
  queue_name              = "terraform-awscc-queue-sse-example"
  sqs_managed_sse_enabled = true
  tags = [{
    key   = "Modified By"
    value = "AWSCC"
  }]
}
