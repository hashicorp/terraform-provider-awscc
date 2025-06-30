---
page_title: "awscc_cloudtrail_trail Resource - terraform-provider-awscc"
subcategory: ""
description: |-
  Creates a trail that specifies the settings for delivery of log data to an Amazon S3 bucket. A maximum of five trails can exist in a region, irrespective of the region in which they were created.
---

# awscc_cloudtrail_trail (Resource)

Creates a trail that specifies the settings for delivery of log data to an Amazon S3 bucket. A maximum of five trails can exist in a region, irrespective of the region in which they were created.

## Example Usage

### Basic Trail

Creates a Cloudtrail with an S3 bucket as the log destination.

```terraform
resource "awscc_cloudtrail_trail" "example" {
  trail_name     = "example"
  is_logging     = true
  s3_bucket_name = awscc_s3_bucket.example.id

  tags = [{
    key   = "Managed By"
    value = "AWSCC"
  }]
}

resource "awscc_s3_bucket_policy" "example" {
  bucket = awscc_s3_bucket.example.id
  policy_document = jsonencode({
    Version = "2012-10-17",
    Statement = [
      {
        Sid    = "AllowSSLRequestsOnly",
        Effect = "Deny",
        Principal = {
          AWS = "*"
        }
        Action = "s3:*",
        Resource = [
          "${awscc_s3_bucket.example.arn}",
          "${awscc_s3_bucket.example.arn}/*"
        ]
        Condition = {
          Bool = {
            "aws:SecureTransport" = "false"
          }
        }
      },
      {
        Sid    = "AWSBucketPermissionsCheck",
        Effect = "Allow",
        Principal = {
          Service = "cloudtrail.amazonaws.com"
        },
        Action   = ["s3:GetBucketAcl", "s3:ListBucket"],
        Resource = "${awscc_s3_bucket.example.arn}"
      },
      {
        Sid    = "AWSCloudTrailWrite",
        Effect = "Allow",
        Principal = {
          Service = "cloudtrail.amazonaws.com"
        },
        Action   = "s3:PutObject",
        Resource = "${awscc_s3_bucket.example.arn}/AWSLogs/${data.aws_caller_identity.current.account_id}/*"
      }
    ]
  })
}

resource "awscc_s3_bucket" "example" {
  bucket_name = "example-cloudtrail-${data.aws_caller_identity.current.account_id}"
}

data "aws_caller_identity" "current" {}
```

### Complex Trail

Creates a Cloudtrail encrypted with a KMS key and advanced event selectors enabled.

```terraform
resource "awscc_cloudtrail_trail" "example" {
  trail_name                    = "example"
  is_logging                    = true
  enable_log_file_validation    = true
  s3_bucket_name                = awscc_s3_bucket.example.id
  s3_key_prefix                 = "prefix"
  include_global_service_events = false
  kms_key_id                    = awscc_kms_key.example.id

  advanced_event_selectors = [{
    name = "Log all S3 objects events"
    field_selectors = [
      {
        field  = "eventCategory"
        equals = ["Data"]
      },
      {
        field  = "resources.type"
        equals = ["AWS::S3::Object"]
      },
      {
        field  = "resources.ARN"
        equals = ["arn:aws:s3:::my-target-bucket/"]
      }
    ]
  }]

  tags = [{
    key   = "Managed By"
    value = "AWSCC"
  }]
}

resource "awscc_s3_bucket_policy" "example" {
  bucket = awscc_s3_bucket.example.id
  policy_document = jsonencode({
    Version = "2012-10-17",
    Statement = [
      {
        Sid    = "AllowSSLRequestsOnly",
        Effect = "Deny",
        Principal = {
          AWS = "*"
        }
        Action = "s3:*",
        Resource = [
          "${awscc_s3_bucket.example.arn}",
          "${awscc_s3_bucket.example.arn}/*"
        ]
        Condition = {
          Bool = {
            "aws:SecureTransport" = "false"
          }
        }
      },
      {
        Sid    = "AWSBucketPermissionsCheck",
        Effect = "Allow",
        Principal = {
          Service = "cloudtrail.amazonaws.com"
        },
        Action   = ["s3:GetBucketAcl", "s3:ListBucket"],
        Resource = "${awscc_s3_bucket.example.arn}"
      },
      {
        Sid    = "AWSCloudTrailWrite",
        Effect = "Allow",
        Principal = {
          Service = "cloudtrail.amazonaws.com"
        },
        Action   = "s3:PutObject",
        Resource = "${awscc_s3_bucket.example.arn}/prefix/AWSLogs/${data.aws_caller_identity.current.account_id}/*"
      }
    ]
  })
}

resource "awscc_s3_bucket" "example" {
  bucket_name = "example-cloudtrail-${data.aws_caller_identity.current.account_id}"

  bucket_encryption = {
    server_side_encryption_configuration = [{
      server_side_encryption_by_default = {
        sse_algorithm     = "aws:kms"
        kms_master_key_id = awscc_kms_key.example.arn
      }
    }]
  }
}

resource "awscc_kms_key" "example" {
  description         = "S3 KMS key"
  enable_key_rotation = true
  key_policy = jsonencode({
    Version = "2012-10-17",
    Statement = [
      {
        Sid    = "Enable IAM User Permissions",
        Effect = "Allow",
        Principal = {
          AWS = "arn:aws:iam::${data.aws_caller_identity.current.account_id}:root"
        },
        Action   = "kms:*",
        Resource = "*"
      },
      {
        Sid    = "Allow CloudTrail to encrypt and decrypt trail",
        Effect = "Allow",
        Principal = {
          Service = "cloudtrail.amazonaws.com"
        },
        Action = [
          "kms:GenerateDataKey*",
          "kms:Decrypt"
        ]
        Resource = "*"
      }
    ]
  })
}

data "aws_caller_identity" "current" {}
```

### Sending Events to CloudWatch Logs

Creates a Cloudtrail that sends events to a CloudWatch log group.

```terraform
resource "awscc_cloudtrail_trail" "example" {
  trail_name                    = "example"
  is_logging                    = true
  s3_bucket_name                = awscc_s3_bucket.example.id
  cloudwatch_logs_log_group_arn = awscc_logs_log_group.example.arn
  cloudwatch_logs_role_arn      = awscc_iam_role.example.arn

  tags = [{
    key   = "Managed By"
    value = "AWSCC"
  }]
}

resource "awscc_logs_log_group" "example" {
  log_group_name = "example"
}

resource "awscc_iam_role" "example" {
  role_name = "cloudtrail_logs_role"
  assume_role_policy_document = jsonencode({
    Version = "2012-10-17"
    Statement = [
      {
        Action = "sts:AssumeRole"
        Effect = "Allow"
        Principal = {
          Service = "cloudtrail.amazonaws.com"
        }
      },
    ]
  })
}
resource "awscc_iam_role_policy" "example" {
  policy_name = "cloudtrail_cloudwatch_logs_policy"
  role_name   = awscc_iam_role.example.role_name
  policy_document = jsonencode({
    Version = "2012-10-17"
    Statement = [
      {
        Action = [
          "logs:CreateLogStream",
          "logs:PutLogEvents"
        ]
        Effect   = "Allow"
        Resource = "arn:aws:logs:${data.aws_region.current.name}:${data.aws_caller_identity.current.account_id}:log-group:${awscc_logs_log_group.example.id}:log-stream:${data.aws_caller_identity.current.account_id}_CloudTrail_${data.aws_region.current.name}*"
      }
    ]
  })
}

data "aws_caller_identity" "current" {}

data "aws_region" "current" {}
```


<!-- schema generated by tfplugindocs -->
## Schema

### Required

- `is_logging` (Boolean) Whether the CloudTrail is currently logging AWS API calls.
- `s3_bucket_name` (String) Specifies the name of the Amazon S3 bucket designated for publishing log files. See Amazon S3 Bucket Naming Requirements.

### Optional

- `advanced_event_selectors` (Attributes Set) The advanced event selectors that were used to select events for the data store. (see [below for nested schema](#nestedatt--advanced_event_selectors))
- `cloudwatch_logs_log_group_arn` (String) Specifies a log group name using an Amazon Resource Name (ARN), a unique identifier that represents the log group to which CloudTrail logs will be delivered. Not required unless you specify CloudWatchLogsRoleArn.
- `cloudwatch_logs_role_arn` (String) Specifies the role for the CloudWatch Logs endpoint to assume to write to a user's log group.
- `enable_log_file_validation` (Boolean) Specifies whether log file validation is enabled. The default is false.
- `event_selectors` (Attributes Set) Use event selectors to further specify the management and data event settings for your trail. By default, trails created without specific event selectors will be configured to log all read and write management events, and no data events. When an event occurs in your account, CloudTrail evaluates the event selector for all trails. For each trail, if the event matches any event selector, the trail processes and logs the event. If the event doesn't match any event selector, the trail doesn't log the event. You can configure up to five event selectors for a trail. (see [below for nested schema](#nestedatt--event_selectors))
- `include_global_service_events` (Boolean) Specifies whether the trail is publishing events from global services such as IAM to the log files.
- `insight_selectors` (Attributes Set) Lets you enable Insights event logging by specifying the Insights selectors that you want to enable on an existing trail. (see [below for nested schema](#nestedatt--insight_selectors))
- `is_multi_region_trail` (Boolean) Specifies whether the trail applies only to the current region or to all regions. The default is false. If the trail exists only in the current region and this value is set to true, shadow trails (replications of the trail) will be created in the other regions. If the trail exists in all regions and this value is set to false, the trail will remain in the region where it was created, and its shadow trails in other regions will be deleted. As a best practice, consider using trails that log events in all regions.
- `is_organization_trail` (Boolean) Specifies whether the trail is created for all accounts in an organization in AWS Organizations, or only for the current AWS account. The default is false, and cannot be true unless the call is made on behalf of an AWS account that is the master account for an organization in AWS Organizations.
- `kms_key_id` (String) Specifies the KMS key ID to use to encrypt the logs delivered by CloudTrail. The value can be an alias name prefixed by 'alias/', a fully specified ARN to an alias, a fully specified ARN to a key, or a globally unique identifier.
- `s3_key_prefix` (String) Specifies the Amazon S3 key prefix that comes after the name of the bucket you have designated for log file delivery. For more information, see Finding Your CloudTrail Log Files. The maximum length is 200 characters.
- `sns_topic_name` (String) Specifies the name of the Amazon SNS topic defined for notification of log file delivery. The maximum length is 256 characters.
- `tags` (Attributes List) (see [below for nested schema](#nestedatt--tags))
- `trail_name` (String)

### Read-Only

- `arn` (String)
- `id` (String) Uniquely identifies the resource.
- `sns_topic_arn` (String)

<a id="nestedatt--advanced_event_selectors"></a>
### Nested Schema for `advanced_event_selectors`

Optional:

- `field_selectors` (Attributes Set) Contains all selector statements in an advanced event selector. (see [below for nested schema](#nestedatt--advanced_event_selectors--field_selectors))
- `name` (String) An optional, descriptive name for an advanced event selector, such as "Log data events for only two S3 buckets".

<a id="nestedatt--advanced_event_selectors--field_selectors"></a>
### Nested Schema for `advanced_event_selectors.field_selectors`

Optional:

- `ends_with` (Set of String) An operator that includes events that match the last few characters of the event record field specified as the value of Field.
- `equals` (Set of String) An operator that includes events that match the exact value of the event record field specified as the value of Field. This is the only valid operator that you can use with the readOnly, eventCategory, and resources.type fields.
- `field` (String) A field in an event record on which to filter events to be logged. Supported fields include readOnly, eventCategory, eventSource (for management events), eventName, resources.type, and resources.ARN.
- `not_ends_with` (Set of String) An operator that excludes events that match the last few characters of the event record field specified as the value of Field.
- `not_equals` (Set of String) An operator that excludes events that match the exact value of the event record field specified as the value of Field.
- `not_starts_with` (Set of String) An operator that excludes events that match the first few characters of the event record field specified as the value of Field.
- `starts_with` (Set of String) An operator that includes events that match the first few characters of the event record field specified as the value of Field.



<a id="nestedatt--event_selectors"></a>
### Nested Schema for `event_selectors`

Optional:

- `data_resources` (Attributes Set) (see [below for nested schema](#nestedatt--event_selectors--data_resources))
- `exclude_management_event_sources` (Set of String) An optional list of service event sources from which you do not want management events to be logged on your trail. In this release, the list can be empty (disables the filter), or it can filter out AWS Key Management Service events by containing "kms.amazonaws.com". By default, ExcludeManagementEventSources is empty, and AWS KMS events are included in events that are logged to your trail.
- `include_management_events` (Boolean) Specify if you want your event selector to include management events for your trail.
- `read_write_type` (String) Specify if you want your trail to log read-only events, write-only events, or all. For example, the EC2 GetConsoleOutput is a read-only API operation and RunInstances is a write-only API operation.

<a id="nestedatt--event_selectors--data_resources"></a>
### Nested Schema for `event_selectors.data_resources`

Optional:

- `type` (String) The resource type in which you want to log data events. You can specify AWS::S3::Object or AWS::Lambda::Function resources.
- `values` (Set of String) An array of Amazon Resource Name (ARN) strings or partial ARN strings for the specified objects.



<a id="nestedatt--insight_selectors"></a>
### Nested Schema for `insight_selectors`

Optional:

- `insight_type` (String) The type of insight to log on a trail.


<a id="nestedatt--tags"></a>
### Nested Schema for `tags`

Optional:

- `key` (String) The key name of the tag. You can specify a value that is 1 to 127 Unicode characters in length and cannot be prefixed with aws:. You can use any of the following characters: the set of Unicode letters, digits, whitespace, _, ., /, =, +, and -.
- `value` (String) The value for the tag. You can specify a value that is 1 to 255 Unicode characters in length and cannot be prefixed with aws:. You can use any of the following characters: the set of Unicode letters, digits, whitespace, _, ., /, =, +, and -.

## Import

Import is supported using the following syntax:

The [`terraform import` command](https://developer.hashicorp.com/terraform/cli/commands/import) can be used, for example:

```shell
$ terraform import awscc_cloudtrail_trail.example "trail_name"
```