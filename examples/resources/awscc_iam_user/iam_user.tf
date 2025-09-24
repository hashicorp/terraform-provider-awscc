resource "awscc_iam_user" "sample_user" {
  user_name = "sample-user"

  login_profile = {
    password                = "MyRandomPassword123$!"
    password_reset_required = true
  }

  policies = [
    {
      policy_name = "sample-policy"
      policy_document = jsonencode({
        "Version" : "2012-10-17",
        "Statement" : [
          {
            "Effect" : "Allow",
            "Action" : [
              "s3:ListAllMyBuckets",
            ],
            "Resource" : "arn:aws:s3:::*"
          }
        ]
      })
    }
  ]

  tags = [
    {
      key   = "Environment"
      value = "Dev"
    },
    {
      key   = "Team"
      value = "DevOps"
    }
  ]
}