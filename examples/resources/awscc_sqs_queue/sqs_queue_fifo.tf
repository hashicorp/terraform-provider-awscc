resource "awscc_sqs_queue" "terraform_awscc_queue_fifo" {
  queue_name                  = "terraform-awscc-queue-example.fifo"
  fifo_queue                  = true
  content_based_deduplication = true
  tags = [{
    key   = "Modified By"
    value = "AWSCC"
  }]

}
