
---
page_title: "awscc_backup_report_plan Resource - terraform-provider-awscc"
subcategory: ""
description: |-
  Contains detailed information about a report plan in AWS Backup Audit Manager.
---

# awscc_backup_report_plan (Resource)

Contains detailed information about a report plan in AWS Backup Audit Manager.

## Example Usage

### AWS Backup Report Plan Configuration

Creates an AWS Backup report plan that generates backup job reports in both CSV and JSON formats, storing them in a specified S3 bucket with custom prefix for a specific AWS account and region.

~> This example is generated by LLM using Amazon Bedrock and validated using terraform validate, apply and destroy. While we strive for accuracy and quality, please note that the information provided may not be entirely error-free or up-to-date. We recommend independently verifying the content.

```terraform
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
```

<!-- schema generated by tfplugindocs -->
## Schema

### Required

- `report_delivery_channel` (Attributes) A structure that contains information about where and how to deliver your reports, specifically your Amazon S3 bucket name, S3 key prefix, and the formats of your reports. (see [below for nested schema](#nestedatt--report_delivery_channel))
- `report_setting` (Attributes) Identifies the report template for the report. Reports are built using a report template. (see [below for nested schema](#nestedatt--report_setting))

### Optional

- `report_plan_description` (String) An optional description of the report plan with a maximum of 1,024 characters.
- `report_plan_name` (String) The unique name of the report plan. The name must be between 1 and 256 characters, starting with a letter, and consisting of letters (a-z, A-Z), numbers (0-9), and underscores (_).
- `report_plan_tags` (Attributes List) Metadata that you can assign to help organize the report plans that you create. Each tag is a key-value pair. (see [below for nested schema](#nestedatt--report_plan_tags))

### Read-Only

- `id` (String) Uniquely identifies the resource.
- `report_plan_arn` (String) An Amazon Resource Name (ARN) that uniquely identifies a resource. The format of the ARN depends on the resource type.

<a id="nestedatt--report_delivery_channel"></a>
### Nested Schema for `report_delivery_channel`

Required:

- `s3_bucket_name` (String) The unique name of the S3 bucket that receives your reports.

Optional:

- `formats` (Set of String) A list of the format of your reports: CSV, JSON, or both. If not specified, the default format is CSV.
- `s3_key_prefix` (String) The prefix for where AWS Backup Audit Manager delivers your reports to Amazon S3. The prefix is this part of the following path: s3://your-bucket-name/prefix/Backup/us-west-2/year/month/day/report-name. If not specified, there is no prefix.


<a id="nestedatt--report_setting"></a>
### Nested Schema for `report_setting`

Required:

- `report_template` (String) Identifies the report template for the report. Reports are built using a report template. The report templates are: `BACKUP_JOB_REPORT | COPY_JOB_REPORT | RESTORE_JOB_REPORT`

Optional:

- `accounts` (Set of String) The list of AWS accounts that a report covers.
- `framework_arns` (Set of String) The Amazon Resource Names (ARNs) of the frameworks a report covers.
- `organization_units` (Set of String) The list of AWS organization units that a report covers.
- `regions` (Set of String) The list of AWS regions that a report covers.


<a id="nestedatt--report_plan_tags"></a>
### Nested Schema for `report_plan_tags`

Optional:

- `key` (String) The key name of the tag. You can specify a value that is 1 to 128 Unicode characters in length and cannot be prefixed with aws:. You can use any of the following characters: the set of Unicode letters, digits, whitespace, _, ., /, =, +, and -.
- `value` (String) The value for the tag. You can specify a value that is 0 to 256 Unicode characters in length and cannot be prefixed with aws:. You can use any of the following characters: the set of Unicode letters, digits, whitespace, _, ., /, =, +, and -.

## Import

Import is supported using the following syntax:

In Terraform v1.5.0 and later, the [`import` block](https://developer.hashicorp.com/terraform/language/import) can be used with the `id` attribute, for example:

```terraform
import {
  to = awscc_backup_report_plan.example
  id = "report_plan_arn"
}
```

The [`terraform import` command](https://developer.hashicorp.com/terraform/cli/commands/import) can be used, for example:

```shell
$ terraform import awscc_backup_report_plan.example "report_plan_arn"
```
