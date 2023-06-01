resource "awscc_sqs_queue" "terraform_awscc_queue_high_throughput" {
  queue_name            = "terraform-awscc-queue-high-throughput-example.fifo"
  fifo_queue            = true
  deduplication_scope   = "messageGroup"
  fifo_throughput_limit = "perMessageGroupId"
  tags = [{
    key   = "Modified By"
    value = "AWSCC"
  }]

}
