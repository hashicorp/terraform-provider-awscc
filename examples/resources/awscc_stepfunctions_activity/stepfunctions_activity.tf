resource "awscc_stepfunctions_activity" "sfn_activity" {
  name = "my-activity"

  tags = [
    {
      key = "Modified By"
      value = "AWSCC"
    }
  ]
}
