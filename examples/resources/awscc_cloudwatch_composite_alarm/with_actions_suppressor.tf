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