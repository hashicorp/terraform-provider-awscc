# Get AWS Account ID and Region
data "aws_caller_identity" "current" {}
data "aws_region" "current" {}

resource "awscc_backup_report_plan" "example" {
  report_plan_name        = "backup_report_example"
  report_plan_description = "Example backup report plan using AWSCC provider"

  report_delivery_channel = {
    s3_bucket_name = "${data.aws_caller_identity.current.account_id}-backup-reports-${data.aws_region.current.name}"
    formats        = ["CSV", "JSON"]
    s3_key_prefix  = "backup-reports"
  }

  report_setting = {
    report_template = "BACKUP_JOB_REPORT"
    regions         = [data.aws_region.current.name]
    accounts        = [data.aws_caller_identity.current.account_id]
  }

  report_plan_tags = [{
    key   = "Modified By"
    value = "AWSCC"
  }]
}