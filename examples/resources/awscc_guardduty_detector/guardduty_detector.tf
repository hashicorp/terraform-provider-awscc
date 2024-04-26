resource "awscc_guardduty_detector" "example" {
  enable = true

  features = [
    {
      name   = "S3_DATA_EVENTS"
      status = "ENABLED"
    },
    {
      name   = "EBS_MALWARE_PROTECTION"
      status = "ENABLED"
    },
    {
      name   = "EKS_AUDIT_LOGS"
      status = "DISABLED"
    }
  ]

  tags = [{
    key   = "Modified By"
    value = "AWSCC"
  }]
} 