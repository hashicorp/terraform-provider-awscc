resource "awscc_iam_managed_policy" "test_policy" {
  description = "Policy for creating a test database"
  path        = "/"

  policy_document = jsonencode(
    {
      "Version" : "2012-10-17",
      "Statement" : [
        {
          "Effect" : "Allow",
          "Action" : "rds:CreateDBInstance",
          "Resource" : "arn:aws:rds:*:*:db:test*",
          "Condition" : {
            "StringEquals" : {
              "rds:DatabaseEngine" : "mysql"
            }
          }
        },
        {
          "Effect" : "Allow",
          "Action" : "rds:CreateDBInstance",
          "Resource" : "arn:aws:rds:*:*:db:test*",
          "Condition" : {
            "StringEquals" : {
              "rds:DatabaseClass" : "db.t2.micro"
            }
          }
        }
      ]
  })

  groups = ["TestDBGroup"]
}