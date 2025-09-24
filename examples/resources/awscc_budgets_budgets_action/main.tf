data "aws_caller_identity" "current" {}

# Create a budget first
resource "aws_budgets_budget" "example" {
  name              = "example-budget"
  budget_type       = "COST"
  limit_amount      = "100"
  limit_unit        = "USD"
  time_unit         = "MONTHLY"
  time_period_start = "2024-01-01_00:00"
}

# Create an IAM role that AWS Budgets can assume
data "aws_iam_policy_document" "budget_assume_role" {
  statement {
    actions = ["sts:AssumeRole"]
    effect  = "Allow"
    principals {
      type        = "Service"
      identifiers = ["budgets.amazonaws.com"]
    }
  }
}

# IAM role for budget action execution
resource "awscc_iam_role" "budget_action" {
  role_name                   = "budget-action-role"
  description                 = "Role for AWS Budget actions"
  assume_role_policy_document = jsonencode(jsondecode(data.aws_iam_policy_document.budget_assume_role.json))
  max_session_duration        = 3600

  tags = [{
    key   = "Modified By"
    value = "AWSCC"
  }]
}

# Policy document for budget actions
data "aws_iam_policy_document" "budget_action_policy" {
  statement {
    effect = "Allow"
    actions = [
      "iam:AttachUserPolicy",
      "iam:DetachUserPolicy"
    ]
    resources = ["arn:aws:iam::${data.aws_caller_identity.current.account_id}:user/*"]
  }
}

# Attach policy to the role
resource "awscc_iam_role_policy" "budget_action" {
  policy_name     = "budget-action-policy"
  role_name       = awscc_iam_role.budget_action.role_name
  policy_document = jsonencode(jsondecode(data.aws_iam_policy_document.budget_action_policy.json))
}

# Create a budget action
resource "awscc_budgets_budgets_action" "example" {
  depends_on         = [aws_budgets_budget.example]
  budget_name        = "example-budget"
  action_type        = "APPLY_IAM_POLICY"
  approval_model     = "AUTOMATIC"
  execution_role_arn = awscc_iam_role.budget_action.arn

  action_threshold = {
    value = 80
    type  = "PERCENTAGE"
  }

  definition = {
    iam_action_definition = {
      policy_arn = "arn:aws:iam::aws:policy/AWSDenyAll"
      users      = ["example-user"]
    }
  }

  notification_type = "ACTUAL"

  subscribers = [
    {
      address = "example@example.com"
      type    = "EMAIL"
    }
  ]

  resource_tags = [{
    key   = "Modified By"
    value = "AWSCC"
  }]
}