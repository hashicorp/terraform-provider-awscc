resource "awscc_connect_routing_profile" "example" {
  name                       = "example"
  description                = "Example routing profile"
  instance_arn               = awscc_connect_instance.example.arn
  default_outbound_queue_arn = awscc_connect_queue.example.queue_arn
  media_concurrencies = [{
    channel     = "TASK"
    concurrency = 1
  }]
  queue_configs = [
    {
      delay    = 0
      priority = 9
      queue_reference = {
        channel   = "TASK"
        queue_arn = awscc_connect_queue.example_task.queue_arn
      }
    }
  ]
  tags = [{
    key   = "ModifiedBy"
    value = "AWSCC"
  }]
}
