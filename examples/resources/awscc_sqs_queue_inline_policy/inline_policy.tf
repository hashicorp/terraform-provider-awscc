resource "awscc_sqs_queue_inline_policy" "example" {
  queue = awscc_sqs_queue.example.queue_url

  policy_document = jsonencode({
    "Version" : "2012-10-17",
    "Statement" : [
      {
        "Effect" : "Allow",
        "Principal" : {
          "AWS" : ["${var.target_account}"]
        },
        "Action" : ["SQS:SendMessage", "SQS:ReceiveMessage"],
        "Resource" : "${awscc_sqs_queue.example.arn}",
      }
    ]
  })
}

resource "awscc_sqs_queue" "example" {
  queue_name = "terraform-awscc-queue-example"
}

variable "target_account" {
  type = string
}