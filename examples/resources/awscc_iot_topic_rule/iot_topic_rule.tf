resource "awscc_iot_topic_rule" "example" {
  rule_name = "example_rule"
  topic_rule_payload = {
    sql         = "SELECT temp FROM 'SomeTopic' WHERE temp > 60"
    description = "example"
    actions = [{
      s3 = {
        bucket_name = "example-bucket"
        key         = "key.txt"
        role_arn    = awscc_iam_role.example.arn
      }
    }]
  }

  tags = [{
    key   = "Modified By"
    value = "AWSCC"
  }]

}