# Create IAM role for the portal
data "aws_iam_policy_document" "portal_assume_role" {
  statement {
    actions = ["sts:AssumeRole"]
    effect  = "Allow"

    principals {
      type = "Service"
      identifiers = [
        "iotsitewise.amazonaws.com",
        "monitor.iotsitewise.amazonaws.com"
      ]
    }
  }
}

# Define portal access policy
data "aws_iam_policy_document" "portal_policy" {
  statement {
    effect = "Allow"
    actions = [
      "iotsitewise:DescribeAsset",
      "iotsitewise:DescribeAssetModel",
      "iotsitewise:DescribeAssetProperty",
      "iotsitewise:GetAssetPropertyValue",
      "iotsitewise:GetAssetPropertyValueHistory",
      "iotsitewise:ListAssets",
      "iotsitewise:ListAssociatedAssets"
    ]
    resources = ["*"]
  }
}

# Create IAM role for portal
resource "awscc_iam_role" "portal_role" {
  role_name                   = "IoTSiteWisePortalRole"
  assume_role_policy_document = jsonencode(jsondecode(data.aws_iam_policy_document.portal_assume_role.json))

  policies = [{
    policy_document = jsonencode(jsondecode(data.aws_iam_policy_document.portal_policy.json))
    policy_name     = "IoTSiteWisePortalAccess"
  }]
}

# Create IoT SiteWise Portal
resource "awscc_iotsitewise_portal" "example" {
  portal_name          = "MyIoTSiteWisePortal"
  portal_description   = "Example IoT SiteWise Portal"
  portal_contact_email = "admin@example.com"
  portal_auth_mode     = "IAM"
  role_arn             = awscc_iam_role.portal_role.arn

  tags = [{
    key   = "Modified By"
    value = "AWSCC"
  }]
}