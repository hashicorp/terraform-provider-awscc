# Create IAM role for sync job
data "aws_iam_policy_document" "assume_role" {
  statement {
    actions = ["sts:AssumeRole"]
    principals {
      type        = "Service"
      identifiers = ["iottwinmaker.amazonaws.com"]
    }
  }
}

data "aws_iam_policy_document" "sync_job_policy" {
  statement {
    effect = "Allow"
    actions = [
      "s3:GetObject",
      "s3:PutObject",
      "s3:ListBucket"
    ]
    resources = ["arn:aws:s3:::example-bucket/*", "arn:aws:s3:::example-bucket"]
  }
  statement {
    effect = "Allow"
    actions = [
      "iottwinmaker:*"
    ]
    resources = ["*"]
  }
}

resource "awscc_iam_role" "sync_job_role" {
  assume_role_policy_document = jsonencode(jsondecode(data.aws_iam_policy_document.assume_role.json))
  description                 = "Role for IoT TwinMaker sync job"
  role_name                   = "IoTTwinMakerSyncJobRole"

  tags = [{
    key   = "Modified By"
    value = "AWSCC"
  }]
}

resource "awscc_iam_role_policy" "sync_job_policy" {
  policy_document = jsonencode(jsondecode(data.aws_iam_policy_document.sync_job_policy.json))
  policy_name     = "IoTTwinMakerSyncJobPolicy"
  role_name       = awscc_iam_role.sync_job_role.role_name
}

# Wait for role propagation
resource "time_sleep" "wait_30_seconds" {
  depends_on      = [awscc_iam_role.sync_job_role]
  create_duration = "30s"
}

# Create IoT TwinMaker workspace (required for sync job)
resource "awscc_iottwinmaker_workspace" "example" {
  depends_on   = [time_sleep.wait_30_seconds]
  workspace_id = "example-workspace"
  role         = awscc_iam_role.sync_job_role.arn
  s3_location  = "arn:aws:s3:::example-bucket"

  tags = {
    "Modified By" = "AWSCC"
  }
}

# Create the sync job
resource "awscc_iottwinmaker_sync_job" "example" {
  workspace_id = awscc_iottwinmaker_workspace.example.workspace_id
  sync_role    = awscc_iam_role.sync_job_role.arn
  sync_source  = "arn:aws:s3:::example-bucket/sync-data"

  tags = {
    "Modified By" = "AWSCC"
  }
}