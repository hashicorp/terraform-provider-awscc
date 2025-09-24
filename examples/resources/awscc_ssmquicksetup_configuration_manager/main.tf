# Get current AWS account ID
data "aws_caller_identity" "current" {}

# Data source for SSM policy document
data "aws_iam_policy_document" "ssm_managed_instance" {
  statement {
    effect = "Allow"
    actions = [
      "ssm:DescribeAssociation",
      "ssm:GetDeployablePatchSnapshotForInstance",
      "ssm:GetDocument",
      "ssm:DescribeDocument",
      "ssm:GetManifest",
      "ssm:GetParameter",
      "ssm:GetParameters",
      "ssm:ListAssociations",
      "ssm:ListInstanceAssociations",
      "ssm:PutInventory",
      "ssm:PutComplianceItems",
      "ssm:PutConfigurePackageResult",
      "ssm:UpdateAssociationStatus",
      "ssm:UpdateInstanceAssociationStatus",
      "ssm:UpdateInstanceInformation"
    ]
    resources = ["*"]
  }

  statement {
    effect = "Allow"
    actions = [
      "ssmmessages:CreateControlChannel",
      "ssmmessages:CreateDataChannel",
      "ssmmessages:OpenControlChannel",
      "ssmmessages:OpenDataChannel"
    ]
    resources = ["*"]
  }

  statement {
    effect = "Allow"
    actions = [
      "ec2messages:AcknowledgeMessage",
      "ec2messages:DeleteMessage",
      "ec2messages:FailMessage",
      "ec2messages:GetEndpoint",
      "ec2messages:GetMessages",
      "ec2messages:SendReply"
    ]
    resources = ["*"]
  }
}

# SSM Quick Setup Configuration Manager
resource "awscc_ssmquicksetup_configuration_manager" "example" {
  configuration_definitions = [{
    name = "SSM-Managed-Instance"
    parameters = {
      "TargetAccounts"     = jsonencode([data.aws_caller_identity.current.account_id])
      "SsmRole"            = "service-role/AWSSystemsManagerFullAccess"
      "ScheduleExpression" = "rate(1 hour)"
    }
    type = "SSM"
  }]

  tags = [{
    key   = "Modified By"
    value = "AWSCC"
  }]
}