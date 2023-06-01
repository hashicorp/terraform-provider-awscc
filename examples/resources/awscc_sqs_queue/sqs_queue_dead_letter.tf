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

resource "awscc_sqs_queue" "terraform_awscc_queue_deadletter" {
  queue_name = "terraform-awscc-queue-deadletter-example"
  redrive_allow_policy = jsonencode({
    redrivePermission = "byQueue",
    sourceQueueArns   = [awscc_sqs_queue.terraform_awscc_queue.arn]
  })
  tags = [{
    key   = "Modified By"
    value = "AWSCC"
  }]

}
