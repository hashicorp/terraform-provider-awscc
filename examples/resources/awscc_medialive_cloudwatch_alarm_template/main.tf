resource "awscc_medialive_cloudwatch_alarm_template_group" "example" {
  name        = "TEST_GROUP"
  description = "Test alarm template group"
  tags = [{
    key   = "Modified By"
    value = "AWSCC"
  }]
}

resource "awscc_medialive_cloudwatch_alarm_template" "example" {
  comparison_operator  = "GreaterThanThreshold"
  group_identifier     = awscc_medialive_cloudwatch_alarm_template_group.example.name
  metric_name          = "NetworkIn"
  name                 = "test-alarm-template"
  statistic            = "Average"
  target_resource_type = "MEDIALIVE_CHANNEL"
  treat_missing_data   = "missing"

  # Optional parameters
  datapoints_to_alarm = 1
  description         = "Example MediaLive CloudWatch Alarm Template"
  evaluation_periods  = 2
  period              = 300
  threshold           = 1000

  tags = [{
    key   = "Modified By"
    value = "AWSCC"
    }, {
    key   = "Environment"
    value = "Test"
  }]
}