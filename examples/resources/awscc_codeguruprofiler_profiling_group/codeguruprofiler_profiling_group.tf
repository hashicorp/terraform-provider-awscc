resource "awscc_codeguruprofiler_profiling_group" "example" {
  profiling_group_name = "example"
  compute_platform     = "Default"

  agent_permissions = {
    principals = [
      var.principal_arn
    ]
  }

  anomaly_detection_notification_configuration = [
    {
      channel_id  = "aaaaaaaa-bbbb-cccc-dddd-eeeeeeeeeeee"
      channel_uri = var.sns_topic_arn
    }
  ]

  tags = [{
    key   = "Modified By"
    value = "AWSCC"
  }]

}

variable "principal_arn" {
  type = string
}

variable "sns_topic_arn" {
  type = string
}
