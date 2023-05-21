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
