resource "awscc_lambda_event_source_mapping" "example" {
  event_source_arn = awscc_sqs_queue.example.arn
  function_name    = awscc_lambda_function.example.arn
  enabled          = true
  batch_size       = 10
}

resource "awscc_sqs_queue" "example" {
  queue_name = "terraform-awscc-queue-example"
}

resource "awscc_lambda_permission" "example" {
  action        = "lambda:InvokeFunction"
  function_name = awscc_lambda_function.example.function_name
  principal     = "sqs.amazonaws.com"
  source_arn    = awscc_sqs_queue.example.arn
}