# First, we need a Connect instance
resource "awscc_connect_instance" "example" {
  identity_management_type = "CONNECT_MANAGED"
  instance_alias           = "example-connect-instance-${formatdate("YYYYMMDDhhmmss", timestamp())}"
  attributes = {
    inbound_calls  = true
    outbound_calls = true
  }
  tags = [{
    key   = "Modified By"
    value = "AWSCC"
  }]
}

# Create a basic queue for our routing profile
resource "awscc_connect_queue" "example_task" {
  instance_arn           = awscc_connect_instance.example.arn
  name                   = "Example-Task-Queue"
  description            = "Example task queue for routing profile"
  hours_of_operation_arn = awscc_connect_hours_of_operation.example.hours_of_operation_arn
  tags = [{
    key   = "Modified By"
    value = "AWSCC"
  }]
}

# Hours of operation are required for queues
resource "awscc_connect_hours_of_operation" "example" {
  instance_arn = awscc_connect_instance.example.arn
  name         = "Example-Hours"
  description  = "Example hours of operation"
  time_zone    = "UTC"
  config = [{
    day = "MONDAY"
    end_time = {
      hours   = 23
      minutes = 59
    }
    start_time = {
      hours   = 0
      minutes = 0
    }
  }]
  tags = [{
    key   = "Modified By"
    value = "AWSCC"
  }]
}

# Create the routing profile
resource "awscc_connect_routing_profile" "example" {
  name                       = "example-routing-profile"
  description                = "Example routing profile created by AWSCC"
  instance_arn               = awscc_connect_instance.example.arn
  default_outbound_queue_arn = awscc_connect_queue.example_task.queue_arn

  media_concurrencies = [{
    channel     = "TASK"
    concurrency = 1
  }]

  queue_configs = [{
    delay    = 0
    priority = 1
    queue_reference = {
      channel   = "TASK"
      queue_arn = awscc_connect_queue.example_task.queue_arn
    }
  }]

  tags = [{
    key   = "Modified By"
    value = "AWSCC"
  }]
}