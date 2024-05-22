data "aws_iam_role" "example" {
  name = "awscc_example"
}

resource "awscc_sqs_queue" "source" {
  fifo_queue = true
}

resource "awscc_sqs_queue" "target" {
  fifo_queue = true
}

resource "awscc_pipes_pipe" "example" {
  name     = "example-pipe"
  role_arn = data.aws_iam_role.example.arn
  source   = awscc_sqs_queue.source.arn
  target   = awscc_sqs_queue.target.arn
  source_parameters = {
    filter_criteria = {
      filters = [{
        pattern = jsonencode({
          source = ["event-source"]
        })
      }]
    }
    sqs_queue_parameters = {
      batch_size = 1
    }
  }
  target_parameters = {
    sqs_queue_parameters = {
      message_deduplication_id = "example-dedupe"
      message_group_id         = "example-group"
    }
  }
}