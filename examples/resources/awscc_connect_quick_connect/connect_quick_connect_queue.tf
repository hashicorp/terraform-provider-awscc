resource "awscc_connect_quick_connect" "example" {
  instance_arn = awscc_connect_instance.example.arn
  name         = "example"
  description  = "example for queue connect type"
  quick_connect_config = {
    quick_connect_type = "QUEUE"
    queue_config = {
      contact_flow_arn = awscc_connect_contact_flow.example.contact_flow_arn
      queue_arn        = awscc_connect_queue.example.queue_arn
    }
  }

  tags = [{
    key   = "ModifiedBy"
    value = "AWSCC"
  }]
}
