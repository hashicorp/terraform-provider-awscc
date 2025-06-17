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


# Create IAM role for the fleet
resource "awscc_iam_role" "complete_fleet_role" {
  role_name = "deadline-fleet-role"
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

  # Add basic permissions for Deadline fleet operations
  managed_policy_arns = [
    "arn:aws:iam::aws:policy/AWSDeadlineCloud-FleetWorker"
  ]
}

# Create the Deadline Fleet
resource "awscc_deadline_fleet" "example" {
  display_name     = "example"
  farm_id          = awscc_deadline_farm.example.farm_id
  max_worker_count = 20
  min_worker_count = 1
  role_arn         = awscc_iam_role.complete_fleet_role.arn

  configuration = {
    service_managed_ec_2 = {
      instance_capabilities = {
        cpu_architecture_type = "x86_64"
        os_family             = "LINUX"
        memory_mi_b = {
          min = 4096
          max = 16384
        }
        v_cpu_count = {
          min = 2
          max = 8
        }
        root_ebs_volume = {
          size_gi_b = 100
        }
      }
      instance_market_options = {
        type = "spot"
      }
    }
  }
}

# Create Queue Fleet Association
resource "awscc_deadline_queue_fleet_association" "complete_association" {
  farm_id  = awscc_deadline_farm.example.farm_id
  queue_id = awscc_deadline_queue.example.queue_id
  fleet_id = awscc_deadline_fleet.example.fleet_id
}
