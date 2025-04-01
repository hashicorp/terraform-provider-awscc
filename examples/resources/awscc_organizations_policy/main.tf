# Define SCP policy document
data "aws_iam_policy_document" "deny_leave_org" {
  statement {
    sid       = "DenyLeaveOrg"
    effect    = "Deny"
    actions   = ["organizations:LeaveOrganization"]
    resources = ["*"]
  }
}

# Create Organizations Policy (SCP)
resource "awscc_organizations_policy" "example" {
  name        = "deny-leave-org"
  description = "Prevents accounts from leaving the organization"
  type        = "SERVICE_CONTROL_POLICY"
  content     = jsonencode(jsondecode(data.aws_iam_policy_document.deny_leave_org.json))

  tags = [{
    key   = "Modified By"
    value = "AWSCC"
  }]
}