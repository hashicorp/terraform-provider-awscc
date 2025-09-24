# Use existing GuardDuty detector
data "aws_guardduty_detector" "main" {}

# Create the GuardDuty filter
resource "awscc_guardduty_filter" "example" {
  name        = "high-severity-findings"
  detector_id = data.aws_guardduty_detector.main.id
  description = "Filter for high severity findings"
  rank        = 1
  action      = "ARCHIVE"

  finding_criteria = {
    criterion = {
      severity = {
        gte = 7
      }
      type = {
        neq = ["UnauthorizedAccess:EC2/SSHBruteForce"]
      }
    }
  }

  tags = [{
    key   = "Modified By"
    value = "AWSCC"
  }]
}