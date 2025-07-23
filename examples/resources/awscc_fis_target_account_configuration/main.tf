data "aws_caller_identity" "current" {}

data "aws_iam_policy_document" "fis_assume_role" {
  statement {
    actions = ["sts:AssumeRole"]
    principals {
      type        = "Service"
      identifiers = ["fis.amazonaws.com"]
    }
  }
}

# Create IAM role for FIS
resource "awscc_iam_role" "fis_role" {
  role_name = "example-target-account-role"
  assume_role_policy_document = jsonencode({
    Version = "2012-10-17"
    Statement = [{
      Effect = "Allow"
      Principal = {
        Service = "fis.amazonaws.com"
      }
      Action = "sts:AssumeRole"
    }]
  })
  description = "IAM role for FIS target account"
  tags = [{
    key   = "Modified By"
    value = "AWSCC"
  }]
}

# Create FIS experiment template
resource "awscc_fis_experiment_template" "example" {
  description = "Example experiment template for target account"
  tags = {
    "Modified By" = "AWSCC"
  }

  actions = {
    "ActionOne" = {
      action_id   = "aws:ec2:stop-instances"
      description = "Stop EC2 instances"
      parameters  = {}
      targets = {
        "Instances" = "ResourceTag"
      }
      action_type = "aws:ec2:stop-instances"
    }
  }
  role_arn = awscc_iam_role.fis_role.arn
  stop_conditions = [{
    source = "none"
  }]
  experiment_options = {
    account_targeting = "multi-account"
  }
  targets = {
    "ResourceTag" = {
      resource_type  = "aws:ec2:instance"
      selection_mode = "ALL"
      filters = [{
        path   = "State.Name"
        values = ["running"]
      }]
      resource_tags = {
        "Environment" = "Test"
      }
    }
  }
}

# Create FIS target account configuration
resource "awscc_fis_target_account_configuration" "example" {
  account_id             = data.aws_caller_identity.current.account_id
  experiment_template_id = awscc_fis_experiment_template.example.id
  role_arn               = awscc_iam_role.fis_role.arn
  description            = "Example FIS target account configuration"
}