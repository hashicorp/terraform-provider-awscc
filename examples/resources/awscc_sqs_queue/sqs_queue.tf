resource "awscc_sqs_queue" "terraform_awscc_queue" {
  queue_name                        = "terraform-awscc-queue-example"
  delay_seconds                     = 90
  maximum_message_size              = 2048
  message_retention_period          = 86400
  receive_message_wait_time_seconds = 10
  tags = [{
    key   = "Modified By"
    value = "AWSCC"
  }]

}
