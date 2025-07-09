# Use AWS standard resource for backup vault since AWSCC version has validation issues
resource "aws_backup_vault" "test" {
  name = "example-backup-vault"
}

resource "awscc_backup_restore_testing_plan" "example" {
  restore_testing_plan_name = "example-restore-testing-plan"
  recovery_point_selection = {
    recovery_point_arn = "${aws_backup_vault.test.arn}/recovery-point/test-recovery-point"
    recovery_point_filters = [
      {
        created_after        = "2024-01-01T00:00:00Z"
        recovery_point_types = ["SNAPSHOT"]
      }
    ]
  }
  schedule_expression = "cron(0 0 ? * 1 *)" # Run weekly on Sundays
}

resource "awscc_backup_restore_testing_selection" "example" {
  restore_testing_selection_name = "example-restore-testing-selection"
  restore_testing_plan_name      = awscc_backup_restore_testing_plan.example.restore_testing_plan_name
  protected_resource_type        = "EC2"
  iam_role_arn                   = aws_iam_role.example.arn
}

# IAM role for AWS Backup restore testing
resource "aws_iam_role" "example" {
  name = "backup-restore-testing-role"

  assume_role_policy = jsonencode({
    Version = "2012-10-17"
    Statement = [
      {
        Action = "sts:AssumeRole"
        Effect = "Allow"
        Principal = {
          Service = "backup.amazonaws.com"
        }
      }
    ]
  })
}

# Custom IAM policy for backup restore testing
resource "aws_iam_role_policy" "example" {
  name = "backup-restore-testing-policy"
  role = aws_iam_role.example.id

  policy = jsonencode({
    Version = "2012-10-17"
    Statement = [
      {
        Effect = "Allow"
        Action = [
          "backup:CreateRestoreTestingPlan",
          "backup:DeleteRestoreTestingPlan",
          "backup:GetRestoreTestingPlan",
          "backup:ListRestoreTestingPlans",
          "backup:StartRestoreTestingPlan",
          "backup:StopRestoreTestingPlan",
          "backup:UpdateRestoreTestingPlan"
        ]
        Resource = "*"
      },
      {
        Effect = "Allow"
        Action = [
          "ec2:CreateImage",
          "ec2:CopyImage",
          "ec2:RunInstances",
          "ec2:StartInstances",
          "ec2:StopInstances",
          "ec2:TerminateInstances",
          "ec2:DeleteSnapshot",
          "ec2:ModifySnapshotAttribute",
          "ec2:ResetSnapshotAttribute",
          "ec2:Describe*"
        ]
        Resource = "*"
      }
    ]
  })
}