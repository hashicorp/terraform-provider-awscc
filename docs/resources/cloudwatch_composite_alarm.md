---
page_title: "awscc_cloudwatch_composite_alarm Resource - terraform-provider-awscc"
subcategory: ""
description: |-
  The AWS::CloudWatch::CompositeAlarm type specifies an alarm which aggregates the states of other Alarms (Metric or Composite Alarms) as defined by the AlarmRule expression
---

# awscc_cloudwatch_composite_alarm (Resource)

The AWS::CloudWatch::CompositeAlarm type specifies an alarm which aggregates the states of other Alarms (Metric or Composite Alarms) as defined by the AlarmRule expression

## Example Usage

### Example with 2 sub-alarms

Creates a Composite alarm that comprises 2 sub-alarms. Note that the AWS provider resource for [aws_cloudwatch_metric_alarm](https://registry.terraform.io/providers/hashicorp/aws/latest/docs/resources/cloudwatch_metric_alarm) is used.

```terraform
resource "awscc_cloudwatch_composite_alarm" "example" {
  alarm_name        = "example-composite-alarm"
  alarm_description = "Example of a composite alarm with various actions"

  alarm_rule = "ALARM(${aws_cloudwatch_metric_alarm.cpu_gte_80.alarm_name}) OR ALARM(${aws_cloudwatch_metric_alarm.status_gte_1.alarm_name})"
}

resource "aws_cloudwatch_metric_alarm" "cpu_gte_80" {
  alarm_name        = "cpu-gte-80"
  alarm_description = "This metric monitors ec2 cpu utilization"

  comparison_operator = "GreaterThanOrEqualToThreshold"
  evaluation_periods  = 2
  metric_name         = "CPUUtilization"
  namespace           = "AWS/EC2"
  period              = 120
  statistic           = "Average"
  threshold           = 80
}

resource "aws_cloudwatch_metric_alarm" "status_gte_1" {
  alarm_name        = "status-gte-1"
  alarm_description = "This metric monitors ec2 status check failed"

  comparison_operator = "GreaterThanOrEqualToThreshold"
  evaluation_periods  = 2
  metric_name         = "StatusCheckFailed"
  namespace           = "AWS/EC2"
  period              = 120
  statistic           = "Average"
  threshold           = 1
}
```

### Example with 2 sub-alarms and various actions

Creates a Composite alarm that comprises 2 sub-alarms. Note that AWS provider resources for [aws_sns_topic](https://registry.terraform.io/providers/hashicorp/aws/latest/docs/resources/sns_topic) and [aws_cloudwatch_metric_alarm](https://registry.terraform.io/providers/hashicorp/aws/latest/docs/resources/cloudwatch_metric_alarm) are used. It also uses different SNS topics for the various alarm actions.

```terraform
resource "awscc_cloudwatch_composite_alarm" "example" {
  alarm_name        = "example-composite-alarm"
  alarm_description = "Example of a composite alarm with various actions"

  alarm_actions             = [aws_sns_topic.example_alarm_actions.arn]
  ok_actions                = [aws_sns_topic.example_ok_actions.arn]
  insufficient_data_actions = [aws_sns_topic.example_insufficient_data_actions.arn]

  alarm_rule = "ALARM(${aws_cloudwatch_metric_alarm.cpu_gte_80.alarm_name}) OR ALARM(${aws_cloudwatch_metric_alarm.status_gte_1.alarm_name})"
}

resource "aws_sns_topic" "example_alarm_actions" {
  name = "example-alarm-actions"
}

resource "aws_sns_topic" "example_ok_actions" {
  name = "example-ok-actions"
}

resource "aws_sns_topic" "example_insufficient_data_actions" {
  name = "example-insufficient-data-actions"
}

resource "aws_cloudwatch_metric_alarm" "cpu_gte_80" {
  alarm_name        = "cpu-gte-80"
  alarm_description = "This metric monitors ec2 cpu utilization"

  comparison_operator = "GreaterThanOrEqualToThreshold"
  evaluation_periods  = 2
  metric_name         = "CPUUtilization"
  namespace           = "AWS/EC2"
  period              = 120
  statistic           = "Average"
  threshold           = 80
}

resource "aws_cloudwatch_metric_alarm" "status_gte_1" {
  alarm_name        = "status-gte-1"
  alarm_description = "This metric monitors ec2 status check failed"

  comparison_operator = "GreaterThanOrEqualToThreshold"
  evaluation_periods  = 2
  metric_name         = "StatusCheckFailed"
  namespace           = "AWS/EC2"
  period              = 120
  statistic           = "Average"
  threshold           = 1
}
```

### Example with actions suppressor

Creates a Composite alarm with an actions suppressor. Note that AWS provider resources for [aws_sns_topic](https://registry.terraform.io/providers/hashicorp/aws/latest/docs/resources/sns_topic) and [aws_cloudwatch_metric_alarm](https://registry.terraform.io/providers/hashicorp/aws/latest/docs/resources/cloudwatch_metric_alarm) are used.

```terraform
resource "awscc_cloudwatch_composite_alarm" "example" {
  alarm_name        = "example-composite-alarm"
  alarm_description = "Example of a composite alarm with actions suppressor"

  actions_suppressor                  = aws_cloudwatch_metric_alarm.cpu_gte_80.alarm_name
  actions_suppressor_extension_period = 60
  actions_suppressor_wait_period      = 60

  alarm_actions = [aws_sns_topic.example_alarm_actions.arn]

  alarm_rule = "ALARM(${aws_cloudwatch_metric_alarm.cpu_gte_80.alarm_name}) OR ALARM(${aws_cloudwatch_metric_alarm.status_gte_1.alarm_name})"
}

resource "aws_sns_topic" "example_alarm_actions" {
  name = "example-alarm-actions"
}

resource "aws_cloudwatch_metric_alarm" "cpu_gte_80" {
  alarm_name                = "cpu-gte-80"
  comparison_operator       = "GreaterThanOrEqualToThreshold"
  evaluation_periods        = 2
  metric_name               = "CPUUtilization"
  namespace                 = "AWS/EC2"
  period                    = 120
  statistic                 = "Average"
  threshold                 = 80
  alarm_description         = "This metric monitors ec2 cpu utilization"
  insufficient_data_actions = []
}

resource "aws_cloudwatch_metric_alarm" "status_gte_1" {
  alarm_name        = "status-gte-1"
  alarm_description = "This metric monitors ec2 status check failed"

  comparison_operator = "GreaterThanOrEqualToThreshold"
  evaluation_periods  = 2
  metric_name         = "StatusCheckFailed"
  namespace           = "AWS/EC2"
  period              = 120
  statistic           = "Average"
  threshold           = 1
}
```

<!-- schema generated by tfplugindocs -->
## Schema

### Required

- `alarm_rule` (String) Expression which aggregates the state of other Alarms (Metric or Composite Alarms)

### Optional

- `actions_enabled` (Boolean) Indicates whether actions should be executed during any changes to the alarm state. The default is TRUE.
- `actions_suppressor` (String) Actions will be suppressed if the suppressor alarm is in the ALARM state. ActionsSuppressor can be an AlarmName or an Amazon Resource Name (ARN) from an existing alarm.
- `actions_suppressor_extension_period` (Number) Actions will be suppressed if WaitPeriod is active. The length of time that actions are suppressed is in seconds.
- `actions_suppressor_wait_period` (Number) Actions will be suppressed if ExtensionPeriod is active. The length of time that actions are suppressed is in seconds.
- `alarm_actions` (List of String) The list of actions to execute when this alarm transitions into an ALARM state from any other state. Specify each action as an Amazon Resource Name (ARN).
- `alarm_description` (String) The description of the alarm
- `alarm_name` (String) The name of the Composite Alarm
- `insufficient_data_actions` (List of String) The actions to execute when this alarm transitions to the INSUFFICIENT_DATA state from any other state. Each action is specified as an Amazon Resource Name (ARN).
- `ok_actions` (List of String) The actions to execute when this alarm transitions to the OK state from any other state. Each action is specified as an Amazon Resource Name (ARN).
- `tags` (Attributes List) A list of key-value pairs to associate with the composite alarm. You can associate as many as 50 tags with an alarm. (see [below for nested schema](#nestedatt--tags))

### Read-Only

- `arn` (String) Amazon Resource Name (ARN) of the alarm
- `id` (String) Uniquely identifies the resource.

<a id="nestedatt--tags"></a>
### Nested Schema for `tags`

Optional:

- `key` (String) A unique identifier for the tag. The combination of tag keys and values can help you organize and categorize your resources.
- `value` (String) The value for the specified tag key.

## Import

Import is supported using the following syntax:

In Terraform v1.5.0 and later, the [`import` block](https://developer.hashicorp.com/terraform/language/import) can be used with the `id` attribute, for example:

```terraform
import {
  to = awscc_cloudwatch_composite_alarm.example
  id = "alarm_name"
}
```

The [`terraform import` command](https://developer.hashicorp.com/terraform/cli/commands/import) can be used, for example:

```shell
$ terraform import awscc_cloudwatch_composite_alarm.example "alarm_name"
```