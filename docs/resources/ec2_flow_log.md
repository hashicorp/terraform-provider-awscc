---
page_title: "awscc_ec2_flow_log Resource - terraform-provider-awscc"
subcategory: ""
description: |-
  Specifies a VPC flow log, which enables you to capture IP traffic for a specific network interface, subnet, or VPC.
---

# awscc_ec2_flow_log (Resource)

Specifies a VPC flow log, which enables you to capture IP traffic for a specific network interface, subnet, or VPC.

## Example Usage

### CloudWatch Loggging

Creates a AWS VPC flow log with CloudWatch Logs as the destination.

```terraform
resource "awscc_ec2_flow_log" "example" {
  deliver_logs_permission_arn = awscc_iam_role.example.arn
  log_destination_type        = "cloud-watch-logs"
  log_destination             = awscc_logs_log_group.example.arn
  traffic_type                = "ALL"
  resource_id                 = var.vpc_id
  resource_type               = "VPC"
  tags = [{
    key   = "Managed By"
    value = "AWSCC"
  }]
}

resource "awscc_logs_log_group" "example" {
  log_group_name = "example"
}

resource "awscc_iam_role" "example" {
  role_name = "cloudwatch_flow_log_role"
  assume_role_policy_document = jsonencode({
    Version = "2012-10-17"
    Statement = [
      {
        Action = "sts:AssumeRole"
        Effect = "Allow"
        Principal = {
          Service = "vpc-flow-logs.amazonaws.com"
        }
      },
    ]
  })
}

resource "awscc_iam_role_policy" "example" {
  policy_name = "example"
  role_name   = awscc_iam_role.example.role_name
  policy_document = jsonencode({
    Version = "2012-10-17"
    Statement = [
      {
        Action = [
          "logs:CreateLogStream",
          "logs:DescribeLogGroups",
          "logs:DescribeLogStreams",
          "logs:PutLogEvents",
        ]
        Effect   = "Allow"
        Resource = "${awscc_logs_log_group.example.arn}:*"
      },
    ]
  })
}
```

### Amazon Data Firehose
Creates a AWS VPC flow log with Amazon Data Firehose as the destination.

```terraform
data "aws_caller_identity" "current" {}

resource "awscc_ec2_flow_log" "example" {
  log_destination      = awscc_kinesisfirehose_delivery_stream.example.arn
  log_destination_type = "kinesis-data-firehose"
  traffic_type         = "ALL"
  resource_id          = var.vpc_id
  resource_type        = "VPC"
  tags = [{
    key   = "Managed By"
    value = "AWSCC"
  }]
}

resource "awscc_kinesisfirehose_delivery_stream" "example" {
  delivery_stream_name = "vpc_flow_log"
  s3_destination_configuration = {
    bucket_arn = awscc_s3_bucket.example.arn
    role_arn   = awscc_iam_role.example.arn
  }
}

resource "awscc_s3_bucket" "example" {
  bucket_name = "example-${data.aws_caller_identity.current.account_id}"
  public_access_block_configuration = {
    block_public_acls       = true
    block_public_policy     = true
    ignore_public_acls      = true
    restrict_public_buckets = true
  }
}

resource "awscc_iam_role" "example" {
  role_name = "firehose_flow_log_role"
  assume_role_policy_document = jsonencode({
    Version = "2012-10-17"
    Statement = [
      {
        Action = "sts:AssumeRole"
        Effect = "Allow"
        Principal = {
          Service = "firehose.amazonaws.com"
        }
      },
    ]
  })
}

resource "awscc_iam_role_policy" "example" {
  policy_name = "example"
  role_name   = awscc_iam_role.example.role_name
  policy_document = jsonencode({
    Version = "2012-10-17"
    Statement = [
      {
        Effect = "Allow",
        Action = [
          "s3:AbortMultipartUpload",
          "s3:GetBucketLocation",
          "s3:GetObject",
          "s3:ListBucket",
          "s3:ListBucketMultipartUploads",
          "s3:PutObject"
        ],
        Resource = [
          "${awscc_s3_bucket.example.arn}",
          "${awscc_s3_bucket.example.arn}/*"
        ]
      },
      {
        Effect = "Allow",
        Action = [
          "kinesis:DescribeStream",
          "kinesis:GetShardIterator",
          "kinesis:GetRecords",
          "kinesis:ListShards"
        ],
        Resource = "${awscc_kinesisfirehose_delivery_stream.example.arn}"
      }
    ]
  })
}
```

### S3 Logging

Creates a AWS VPC flow log with S3 as the destination.

```terraform
data "aws_caller_identity" "current" {}

resource "awscc_ec2_flow_log" "example" {
  log_destination      = awscc_s3_bucket.example.arn
  log_destination_type = "s3"
  traffic_type         = "ALL"
  resource_id          = var.vpc_id
  resource_type        = "VPC"
  tags = [{
    key   = "Managed By"
    value = "AWSCC"
  }]
}

resource "awscc_s3_bucket" "example" {
  bucket_name = "example-${data.aws_caller_identity.current.account_id}"
  public_access_block_configuration = {
    block_public_acls       = true
    block_public_policy     = true
    ignore_public_acls      = true
    restrict_public_buckets = true
  }
}
```

### S3 Logging in Parquet Format

Creates a AWS VPC flow log with S3 as the destination in Parquet file format.

```terraform
data "aws_caller_identity" "current" {}

resource "awscc_ec2_flow_log" "example" {
  log_destination      = awscc_s3_bucket.example.arn
  log_destination_type = "s3"
  traffic_type         = "ALL"
  resource_id          = var.vpc_id
  resource_type        = "VPC"
  destination_options = {
    file_format                = "parquet"
    per_hour_partition         = true
    hive_compatible_partitions = true
  }
  tags = [{
    key   = "Managed By"
    value = "AWSCC"
  }]
}

resource "awscc_s3_bucket" "example" {
  bucket_name = "example-${data.aws_caller_identity.current.account_id}-p"
  public_access_block_configuration = {
    block_public_acls       = true
    block_public_policy     = true
    ignore_public_acls      = true
    restrict_public_buckets = true
  }
}
```

<!-- schema generated by tfplugindocs -->
## Schema

### Required

- `resource_id` (String) The ID of the subnet, network interface, or VPC for which you want to create a flow log.
- `resource_type` (String) The type of resource for which to create the flow log. For example, if you specified a VPC ID for the ResourceId property, specify VPC for this property.

### Optional

- `deliver_cross_account_role` (String) The ARN of the IAM role that allows Amazon EC2 to publish flow logs across accounts.
- `deliver_logs_permission_arn` (String) The ARN for the IAM role that permits Amazon EC2 to publish flow logs to a CloudWatch Logs log group in your account. If you specify LogDestinationType as s3 or kinesis-data-firehose, do not specify DeliverLogsPermissionArn or LogGroupName.
- `destination_options` (Attributes) (see [below for nested schema](#nestedatt--destination_options))
- `log_destination` (String) Specifies the destination to which the flow log data is to be published. Flow log data can be published to a CloudWatch Logs log group, an Amazon S3 bucket, or a Kinesis Firehose stream. The value specified for this parameter depends on the value specified for LogDestinationType.
- `log_destination_type` (String) Specifies the type of destination to which the flow log data is to be published. Flow log data can be published to CloudWatch Logs or Amazon S3.
- `log_format` (String) The fields to include in the flow log record, in the order in which they should appear.
- `log_group_name` (String) The name of a new or existing CloudWatch Logs log group where Amazon EC2 publishes your flow logs. If you specify LogDestinationType as s3 or kinesis-data-firehose, do not specify DeliverLogsPermissionArn or LogGroupName.
- `max_aggregation_interval` (Number) The maximum interval of time during which a flow of packets is captured and aggregated into a flow log record. You can specify 60 seconds (1 minute) or 600 seconds (10 minutes).
- `tags` (Attributes List) The tags to apply to the flow logs. (see [below for nested schema](#nestedatt--tags))
- `traffic_type` (String) The type of traffic to log. You can log traffic that the resource accepts or rejects, or all traffic.

### Read-Only

- `flow_log_id` (String) The Flow Log ID
- `id` (String) Uniquely identifies the resource.

<a id="nestedatt--destination_options"></a>
### Nested Schema for `destination_options`

Optional:

- `file_format` (String)
- `hive_compatible_partitions` (Boolean)
- `per_hour_partition` (Boolean)


<a id="nestedatt--tags"></a>
### Nested Schema for `tags`

Optional:

- `key` (String)
- `value` (String)

## Import

Import is supported using the following syntax:

The [`terraform import` command](https://developer.hashicorp.com/terraform/cli/commands/import) can be used, for example:

```shell
$ terraform import awscc_ec2_flow_log.example "id"
```