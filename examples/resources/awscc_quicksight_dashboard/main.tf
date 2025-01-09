# Get current AWS account ID
data "aws_caller_identity" "current" {}

# Get current AWS region
data "aws_region" "current" {}

# QuickSight dashboard permissions policy
data "aws_iam_policy_document" "dashboard_permissions" {
  statement {
    actions = [
      "quicksight:DescribeDashboard",
      "quicksight:ListDashboardVersions",
      "quicksight:QueryDashboard"
    ]
    resources = ["*"]
    effect    = "Allow"
  }
}

# Create the QuickSight dashboard
resource "awscc_quicksight_dashboard" "example" {
  aws_account_id = data.aws_caller_identity.current.account_id
  dashboard_id   = "example-dashboard"
  name           = "Example Dashboard"

  source_entity = {
    source_template = {
      arn = "arn:aws:quicksight:${data.aws_region.current.name}:${data.aws_caller_identity.current.account_id}:template/example-template"
      data_set_references = [
        {
          data_set_arn         = "arn:aws:quicksight:${data.aws_region.current.name}:${data.aws_caller_identity.current.account_id}:dataset/example-dataset"
          data_set_placeholder = "example_dataset"
        }
      ]
    }
  }

  permissions = [
    {
      principal = "arn:aws:quicksight:${data.aws_region.current.name}:${data.aws_caller_identity.current.account_id}:user/default/admin"
      actions = [
        "quicksight:DescribeDashboard",
        "quicksight:ListDashboardVersions",
        "quicksight:QueryDashboard",
        "quicksight:UpdateDashboardPermissions",
        "quicksight:UpdateDashboard"
      ]
    }
  ]

  dashboard_publish_options = {
    ad_hoc_filtering_option = {
      availability_status = "ENABLED"
    }
    export_to_csv_option = {
      availability_status = "ENABLED"
    }
    sheet_controls_option = {
      visibility_state = "EXPANDED"
    }
  }

  version_description = "Initial version of the dashboard"

  tags = [{
    key   = "Modified By"
    value = "AWSCC"
  }]
}