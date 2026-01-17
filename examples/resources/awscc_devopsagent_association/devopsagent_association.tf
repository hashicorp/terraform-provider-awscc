resource "awscc_devopsagent_association" "example" {
  agent_space_id = "c10acee3-02f9-4729-9637-39e0aff6198b"
  service_id     = "codecommit"
  configuration = {
    source_aws = {
      repository_arn     = "arn:aws:codecommit:us-east-1:697621333100:repository/example-repo"
      account_id         = "697621333100"
      account_type       = "source"
      assumable_role_arn = "arn:aws:iam::697621333100:role/devops-agent-role"
    }
  }
}
