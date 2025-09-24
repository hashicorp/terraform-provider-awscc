data "aws_iam_policy_document" "iotsitewise_portal_policy" {
  statement {
    effect = "Allow"
    actions = [
      "iotsitewise:DescribePortal",
      "iotsitewise:ListProjects",
      "iotsitewise:CreateProject",
      "iotsitewise:DeleteProject"
    ]
    resources = ["*"]
  }
}

# Create a service role for IoT SiteWise Portal
resource "awscc_iam_role" "iotsitewise_portal_role" {
  role_name = "IoTSiteWisePortalRole"
  assume_role_policy_document = jsonencode({
    Version = "2012-10-17"
    Statement = [
      {
        Action = "sts:AssumeRole"
        Effect = "Allow"
        Principal = {
          Service = "monitor.iotsitewise.amazonaws.com"
        }
      }
    ]
  })
  tags = [{
    key   = "Modified By"
    value = "AWSCC"
  }]
}

# Attach policy to the role
resource "awscc_iam_role_policy" "iotsitewise_portal_policy" {
  policy_document = data.aws_iam_policy_document.iotsitewise_portal_policy.json
  policy_name     = "IoTSiteWisePortalPolicy"
  role_name       = awscc_iam_role.iotsitewise_portal_role.role_name
}

# Create IoT SiteWise Portal
resource "awscc_iotsitewise_portal" "example" {
  portal_name        = "ExamplePortal"
  portal_description = "Example IoT SiteWise Portal"
  role_arn           = awscc_iam_role.iotsitewise_portal_role.arn

  portal_contact_email = "example@example.com"

  tags = [{
    key   = "Modified By"
    value = "AWSCC"
  }]
}

# Create IoT SiteWise Project
resource "awscc_iotsitewise_project" "example" {
  portal_id           = awscc_iotsitewise_portal.example.portal_id
  project_name        = "ExampleProject"
  project_description = "Example IoT SiteWise Project"

  tags = [{
    key   = "Modified By"
    value = "AWSCC"
  }]
}