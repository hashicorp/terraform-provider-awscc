resource "awscc_datazone_domain" "example" {
  name                  = "example"
  domain_execution_role = awscc_iam_role.example.arn
  description           = "Datazone domain example"


  tags = [{
    key   = "Modified By"
    value = "AWSCC"
  }]
}

resource "awscc_iam_role" "example" {
  path = "/service-role/"
  assume_role_policy_document = jsonencode({
    "Version" : "2012-10-17",
    "Statement" : [
      {
        "Effect" : "Allow",
        "Principal" : {
          "Service" : "datazone.amazonaws.com"
        },
        "Action" : [
          "sts:AssumeRole",
          "sts:TagSession"
        ],
        "Condition" : {
          "StringEquals" : {
            "aws:SourceAccount" : var.source_account_id
          },
          "ForAllValues:StringLike" : {
            "aws:TagKeys" : "datazone*"
          }
        }
      }
    ]
  })
  managed_policy_arns = ["arn:aws:iam::aws:policy/service-role/AmazonDataZoneDomainExecutionRolePolicy"]
}

variable "source_account_id" {
  type        = string
  description = "Source AWS account id"
}