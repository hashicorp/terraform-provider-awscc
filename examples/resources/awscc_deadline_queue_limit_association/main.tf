resource "awscc_deadline_farm" "example" {
  display_name = "example"
  description  = "Example"
  tags = [{
    key   = "ModifiedBy"
    value = "AWSCC"
  }]
}

# Create IAM role for the queue session
resource "awscc_iam_role" "queue_session_role" {
  role_name = "example"
  assume_role_policy_document = jsonencode({
    Version = "2012-10-17"
    Statement = [
      {
        Action = "sts:AssumeRole"
        Effect = "Allow"
        Principal = {
          Service = "credentials.deadline.amazonaws.com"
        }
      }
    ]
  })

  # Add basic permissions for queue session operations
  managed_policy_arns = [
    "arn:aws:iam::aws:policy/AWSDeadlineCloud-UserAccessJobs"
  ]
}

# Create the Deadline Queue
resource "awscc_deadline_queue" "example" {
  display_name = "example"
  farm_id      = awscc_deadline_farm.example.farm_id

  job_run_as_user = {
    run_as = "QUEUE_CONFIGURED_USER"
    posix = {
      user  = "deadline-user"
      group = "deadline-group"
    }
  }

  role_arn = awscc_iam_role.queue_session_role.arn
}

resource "awscc_deadline_limit" "example" {
  display_name            = "CPULimit"
  farm_id                 = awscc_deadline_farm.example.farm_id
  amount_requirement_name = "amount.cpu_cores"
  max_count               = 100
}


resource "awscc_deadline_queue_limit_association" "cpu_association" {
  farm_id  = awscc_deadline_farm.example.farm_id
  queue_id = awscc_deadline_queue.example.queue_id
  limit_id = awscc_deadline_limit.example.limit_id
}
