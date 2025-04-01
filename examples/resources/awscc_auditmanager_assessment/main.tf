# Get current AWS account details
data "aws_caller_identity" "current" {}

# IAM role for Audit Manager service
data "aws_iam_policy_document" "audit_manager_trust" {
  statement {
    effect = "Allow"
    principals {
      type        = "Service"
      identifiers = ["auditmanager.amazonaws.com"]
    }
    actions = ["sts:AssumeRole"]
  }
}

data "aws_iam_policy_document" "audit_manager_policy" {
  statement {
    effect = "Allow"
    actions = [
      "s3:GetObject",
      "s3:ListBucket",
      "s3:PutObject"
    ]
    resources = [
      aws_s3_bucket.audit_reports.arn,
      "${aws_s3_bucket.audit_reports.arn}/*"
    ]
  }
}

resource "aws_iam_role" "audit_manager" {
  name               = "AuditManagerRole"
  assume_role_policy = data.aws_iam_policy_document.audit_manager_trust.json
}

resource "aws_iam_role_policy" "audit_manager" {
  name   = "AuditManagerPolicy"
  role   = aws_iam_role.audit_manager.id
  policy = data.aws_iam_policy_document.audit_manager_policy.json
}

# S3 bucket for assessment reports
resource "aws_s3_bucket" "audit_reports" {
  bucket = "audit-manager-reports-${data.aws_caller_identity.current.account_id}"
}

# Example Audit Manager Assessment
resource "awscc_auditmanager_assessment" "example" {
  name        = "Example-Assessment"
  description = "Example assessment created using AWSCC provider"

  framework_id = "d2e1fd0c-1f31-4084-955a-97f57b0292ea" # Standard AWS Audit Manager Security Best Practices v1.0 framework

  assessment_reports_destination = {
    destination      = "s3://${aws_s3_bucket.audit_reports.id}"
    destination_type = "S3"
  }

  aws_account = {
    id            = data.aws_caller_identity.current.account_id
    email_address = "example@example.com"
    name          = "Example Account"
  }

  roles = [{
    role_arn  = aws_iam_role.audit_manager.arn
    role_type = "PROCESS_OWNER"
  }]

  scope = {
    aws_accounts = [{
      id            = data.aws_caller_identity.current.account_id
      email_address = "example@example.com"
      name          = "Example Account"
    }]
    aws_services = [{
      service_name = "Amazon S3"
    }]
  }

  tags = [{
    key   = "ModifiedBy"
    value = "AWSCC"
  }]
}