# Data source for assume role policy document
data "aws_iam_policy_document" "fleet_assume_role" {
  statement {
    actions = ["sts:AssumeRole"]
    principals {
      type        = "Service"
      identifiers = ["deadline.amazonaws.com"]
    }
  }
}

# IAM role for Deadline fleet using AWSCC
resource "awscc_iam_role" "fleet_role" {
  role_name                   = "deadline-fleet-role"
  assume_role_policy_document = data.aws_iam_policy_document.fleet_assume_role.json

  tags = [{
    key   = "Modified By"
    value = "AWSCC"
  }]
}

# Basic policy for fleet operations
data "aws_iam_policy_document" "fleet_policy" {
  statement {
    actions = [
      "ec2:DescribeInstances",
      "ec2:RunInstances",
      "ec2:TerminateInstances",
      "iam:PassRole"
    ]
    resources = ["*"]
  }
}

# Attach policy to role using AWSCC
resource "awscc_iam_role_policy" "fleet_policy" {
  policy_name     = "deadline-fleet-policy"
  role_name       = awscc_iam_role.fleet_role.role_name
  policy_document = data.aws_iam_policy_document.fleet_policy.json
}

# Deadline fleet resource
resource "awscc_deadline_fleet" "example" {
  display_name     = "example-fleet"
  farm_id          = "farm-12345678901234567890123456789012" # Replace with your actual farm ID
  max_worker_count = 10
  min_worker_count = 1
  role_arn         = awscc_iam_role.fleet_role.arn
  description      = "Example Deadline Fleet using AWSCC provider"

  configuration = {
    service_managed_ec_2 = {
      instance_capabilities = {
        allowed_instance_types = ["t3.medium", "t3.large"]
        cpu_architecture_type  = "x86_64"
        os_family              = "LINUX"
        memory_mi_b = {
          min = 4096 # 4GB
          max = 8192 # 8GB
        }
        v_cpu_count = {
          min = 2
          max = 4
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

  tags = [{
    key   = "Modified By"
    value = "AWSCC"
  }]
}