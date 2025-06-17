data "aws_caller_identity" "current" {}
data "aws_region" "current" {}

resource "awscc_quicksight_folder" "example" {
  aws_account_id = data.aws_caller_identity.current.account_id
  folder_id      = "analytics-team-folder"
  name           = "example"
  folder_type    = "SHARED"
  sharing_model  = "ACCOUNT"

  # Grant permissions to users and groups
  permissions = [
    {
      principal = "arn:aws:quicksight:${data.aws_region.current.name}:${data.aws_caller_identity.current.account_id}:user/default/analytics-admin"
      actions = [
        "quicksight:CreateFolder",
        "quicksight:DescribeFolder",
        "quicksight:UpdateFolder",
        "quicksight:DeleteFolder",
        "quicksight:CreateFolderMembership",
        "quicksight:DeleteFolderMembership",
        "quicksight:DescribeFolderPermissions",
        "quicksight:UpdateFolderPermissions"
      ]
    },
    {
      principal = "arn:aws:quicksight:${data.aws_region.current.name}:${data.aws_caller_identity.current.account_id}:group/default/analytics-team"
      actions = [
        "quicksight:DescribeFolder",
        "quicksight:CreateFolderMembership"
      ]
    }
  ]

  tags = [
    {
      key   = "ModifiedBy"
      value = "AWSCC"
    }
  ]
}
