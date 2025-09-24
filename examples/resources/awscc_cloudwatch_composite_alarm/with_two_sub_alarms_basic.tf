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