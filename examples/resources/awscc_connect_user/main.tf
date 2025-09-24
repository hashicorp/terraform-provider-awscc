data "aws_caller_identity" "current" {}
data "aws_region" "current" {}

# Generate a random suffix for unique usernames
resource "random_id" "user_suffix" {
  byte_length = 4
}

# Create Amazon Connect instance
resource "awscc_connect_instance" "example" {
  identity_management_type = "CONNECT_MANAGED"
  instance_alias           = "example-awscc-connect-instance"
  attributes = {
    inbound_calls    = true
    outbound_calls   = true
    contact_lens     = true
    early_media      = true
    contactflow_logs = true
  }

  tags = [{
    key   = "Modified By"
    value = "AWSCC"
  }]
}

# Create hours of operation (24/7)
resource "awscc_connect_hours_of_operation" "example" {
  instance_arn = awscc_connect_instance.example.arn
  name         = "24x7-Hours"
  description  = "24 hours a day, 7 days a week operation hours"
  time_zone    = "UTC"

  config = [
    {
      day = "MONDAY"
      start_time = {
        hours   = 0
        minutes = 0
      }
      end_time = {
        hours   = 23
        minutes = 59
      }
    },
    {
      day = "TUESDAY"
      start_time = {
        hours   = 0
        minutes = 0
      }
      end_time = {
        hours   = 23
        minutes = 59
      }
    },
    {
      day = "WEDNESDAY"
      start_time = {
        hours   = 0
        minutes = 0
      }
      end_time = {
        hours   = 23
        minutes = 59
      }
    },
    {
      day = "THURSDAY"
      start_time = {
        hours   = 0
        minutes = 0
      }
      end_time = {
        hours   = 23
        minutes = 59
      }
    },
    {
      day = "FRIDAY"
      start_time = {
        hours   = 0
        minutes = 0
      }
      end_time = {
        hours   = 23
        minutes = 59
      }
    },
    {
      day = "SATURDAY"
      start_time = {
        hours   = 0
        minutes = 0
      }
      end_time = {
        hours   = 23
        minutes = 59
      }
    },
    {
      day = "SUNDAY"
      start_time = {
        hours   = 0
        minutes = 0
      }
      end_time = {
        hours   = 23
        minutes = 59
      }
    }
  ]

  tags = [{
    key   = "ModifiedBy"
    value = "AWSCC"
  }]
}

# Create a queue for the routing profile
resource "awscc_connect_queue" "example" {
  instance_arn           = awscc_connect_instance.example.arn
  name                   = "BasicQueueAWSCC"
  description            = "Basic queue for customer support"
  hours_of_operation_arn = awscc_connect_hours_of_operation.example.hours_of_operation_arn

  tags = [{
    key   = "ModifiedBy"
    value = "AWSCC"
  }]
}

# Create routing profile
resource "awscc_connect_routing_profile" "example" {
  instance_arn               = awscc_connect_instance.example.arn
  name                       = "BasicAgentRoutingProfile"
  description                = "Basic routing profile for customer service agents"
  default_outbound_queue_arn = awscc_connect_queue.example.queue_arn

  media_concurrencies = [
    {
      channel     = "VOICE"
      concurrency = 1
    },
    {
      channel     = "CHAT"
      concurrency = 2
    },
    {
      channel     = "TASK"
      concurrency = 3
    }
  ]

  queue_configs = [
    {
      queue_reference = {
        queue_arn = awscc_connect_queue.example.queue_arn
        channel   = "VOICE"
      }
      priority = 1
      delay    = 0
    }
  ]

  tags = [{
    key   = "ModifiedBy"
    value = "AWSCC"
  }]
}

# Create security profile with agent permissions
resource "awscc_connect_security_profile" "example" {
  instance_arn          = awscc_connect_instance.example.arn
  security_profile_name = "AgentSecurityProfile"
  description           = "Security profile for customer service agents"

  permissions = [
    "BasicAgentAccess",
    "OutboundCallAccess"
  ]

  tags = [{
    key   = "ModifiedBy"
    value = "AWSCC"
  }]
}

resource "awscc_connect_user" "example" {
  instance_arn          = awscc_connect_instance.example.arn
  username              = "john.doe.${random_id.user_suffix.hex}"
  password              = "TempPassword123!"
  routing_profile_arn   = awscc_connect_routing_profile.example.routing_profile_arn
  security_profile_arns = [awscc_connect_security_profile.example.security_profile_arn]

  identity_info = {
    first_name      = "John"
    last_name       = "Doe"
    email           = "john.doe+${random_id.user_suffix.hex}@example.com"
    mobile          = "+12345678901"
    secondary_email = "john.doe.backup+${random_id.user_suffix.hex}@example.com"
  }

  phone_config = {
    phone_type                    = "SOFT_PHONE"
    auto_accept                   = true
    after_contact_work_time_limit = 120
  }

  tags = [{
    key   = "ModifiedBy"
    value = "AWSCC"
  }]
}
