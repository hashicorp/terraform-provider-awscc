resource "aws_cloudwatch_metric_alarm" "example_alarm" {
  alarm_name          = "example-high-cpu-alarm"
  comparison_operator = "GreaterThanThreshold"
  evaluation_periods  = "2"
  metric_name         = "CPUUtilization"
  namespace           = "AWS/EC2"
  period              = "120"
  statistic           = "Average"
  threshold           = "80"
  alarm_description   = "This metric monitors ec2 cpu utilization"

  dimensions = {
    InstanceId = "i-1234567890abcdef0"
  }
}

resource "awscc_cloudwatch_alarm_mute_rule" "example" {
  name        = "example-alarm-mute-rule"
  description = "Example alarm mute rule for scheduled maintenance windows"

  rule = {
    schedule = {
      expression = "cron(0 3 * * MON)"
      duration   = "PT1H"
    }
  }

  mute_targets = {
    alarm_names = [aws_cloudwatch_metric_alarm.example_alarm.alarm_name]
  }

  tags = [
    {
      key   = "Name"
      value = "example-alarm-mute-rule"
    }
  ]

  depends_on = [aws_cloudwatch_metric_alarm.example_alarm]
}
