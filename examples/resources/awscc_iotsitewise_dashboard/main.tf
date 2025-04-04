# IAM Role for IoT SiteWise Portal
resource "awscc_iam_role" "portal" {
  role_name = "AWSIoTSiteWiseMonitorPortalAccess"

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

resource "aws_iam_role_policy_attachment" "portal" {
  role       = awscc_iam_role.portal.role_name
  policy_arn = "arn:aws:iam::aws:policy/service-role/AWSIoTSiteWiseMonitorPortalAccess"
}

resource "awscc_iotsitewise_portal" "example" {
  portal_name          = "ExamplePortal"
  portal_contact_email = "example@example.com"
  portal_description   = "Example IoT SiteWise Portal for Dashboard"

  role_arn = awscc_iam_role.portal.arn

  tags = [{
    key   = "Modified By"
    value = "AWSCC"
  }]
}

resource "awscc_iotsitewise_project" "example" {
  project_name        = "ExampleProject"
  project_description = "Example IoT SiteWise Project for Dashboard"
  portal_id           = awscc_iotsitewise_portal.example.portal_id

  tags = [{
    key   = "Modified By"
    value = "AWSCC"
  }]
}

resource "awscc_iotsitewise_dashboard" "example" {
  dashboard_name        = "ExampleDashboard"
  dashboard_description = "Example IoT SiteWise Dashboard created with AWSCC provider"
  project_id            = awscc_iotsitewise_project.example.project_id
  dashboard_definition = jsonencode({
    widgets = [
      {
        type    = "monitor",
        title   = "Example Widget",
        x       = 0,
        y       = 0,
        width   = 4,
        height  = 6,
        metrics = []
      }
    ]
  })

  tags = [{
    key   = "Modified By"
    value = "AWSCC"
  }]
}