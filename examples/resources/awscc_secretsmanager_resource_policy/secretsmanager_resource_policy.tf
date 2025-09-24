resource "awscc_secretsmanager_resource_policy" "example" {

  resource_policy = jsonencode({
    "Version" : "2012-10-17",
    "Statement" : [
      {
        "Effect" : "Deny",
        "Principal" : {
          "AWS" : "arn:aws:iam::${data.aws_caller_identity.current.account_id}:root"
        }
        "Action" : "secretsmanager:DeleteSecret",
        "Resource" : "*"
      }
    ]
  })
  secret_id           = awscc_secretsmanager_secret.example.id
  block_public_policy = true
}

resource "awscc_secretsmanager_secret" "example" {
  name        = "example"
  description = "example"
}

data "aws_caller_identity" "current" {}