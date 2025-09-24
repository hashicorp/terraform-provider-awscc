resource "awscc_cloudwatch_alarm" "lambda_invocations_alarm" {
  alarm_name          = "LambdaInvocationsAlarm"
  comparison_operator = "LessThanLowerOrGreaterThanUpperThreshold"
  evaluation_periods  = 1

  metrics = [{
    expression = "ANOMALY_DETECTION_BAND(m1, 2)"
    id         = "ad1"
    },
    {
      id = "m1"
      metric_stat = {
        metric = {
          metric_name = "Invocations"
          namespace   = "AWS/Lambda"
        }
        period = 86400
        stat   = "Sum"
      }
  }]

  threshold_metric_id = "ad1"
  treat_missing_data  = "breaching"
}