# Get current AWS region and account ID
data "aws_region" "current" {}
data "aws_caller_identity" "current" {}

# Create an IAM role for Redshift scheduled action
resource "awscc_iam_role" "redshift_scheduled_action_role" {
  role_name = "RedshiftScheduledActionRole"
  assume_role_policy_document = jsonencode({
    Version = "2012-10-17"
    Statement = [
      {
        Action = "sts:AssumeRole"
        Effect = "Allow"
        Principal = {
          Service = "scheduler.redshift.amazonaws.com"
        }
      }
    ]
  })
  tags = [{
    key   = "Modified By"
    value = "AWSCC"
  }]
}

# Create IAM policy for Redshift scheduled action
data "aws_iam_policy_document" "redshift_scheduled_action" {
  statement {
    effect = "Allow"
    actions = [
      "redshift:PauseCluster",
      "redshift:ResumeCluster",
      "redshift:ResizeCluster"
    ]
    resources = [
      "arn:aws:redshift:${data.aws_region.current.name}:${data.aws_caller_identity.current.account_id}:cluster:*"
    ]
  }
}

resource "awscc_iam_role_policy" "redshift_scheduled_action_policy" {
  policy_document = jsonencode(jsondecode(data.aws_iam_policy_document.redshift_scheduled_action.json))
  policy_name     = "RedshiftScheduledActionPolicy"
  role_name       = awscc_iam_role.redshift_scheduled_action_role.role_name
}

# Create a Redshift Scheduled Action
resource "awscc_redshift_scheduled_action" "example" {
  scheduled_action_name = "pause-cluster-at-night"
  enable                = true
  schedule              = "cron(0 0 * * ? *)" # Run at midnight UTC every day
  iam_role              = "arn:aws:iam::${data.aws_caller_identity.current.account_id}:role/RedshiftScheduledActionRole"

  target_action = {
    pause_cluster = {
      cluster_identifier = "my-redshift-cluster"
    }
  }

  scheduled_action_description = "Pause the Redshift cluster at midnight UTC"

  depends_on = [awscc_iam_role_policy.redshift_scheduled_action_policy]
}