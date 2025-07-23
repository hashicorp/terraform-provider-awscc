data "aws_caller_identity" "current" {}

resource "awscc_logs_log_group" "example" {
  log_group_name = "my-log-group"

  data_protection_policy = jsonencode({
    "Name" : "data-protection-policy",
    "Description" : "test description",
    "Version" : "2021-06-01",
    "Statement" : [
      {
        "Sid" : "audit-policy test",
        "DataIdentifier" : [
          "arn:aws:dataprotection::aws:data-identifier/EmailAddress",
          "arn:aws:dataprotection::aws:data-identifier/DriversLicense-US"
        ],
        "Operation" : {
          "Audit" : {
            "FindingsDestination" : {
              "CloudWatchLogs" : {
                "LogGroup" : "${awscc_logs_log_group.finding.id}"
              }
            }
          }
        }
      },
      {
        "Sid" : "redact-policy",
        "DataIdentifier" : [
          "arn:aws:dataprotection::aws:data-identifier/EmailAddress",
          "arn:aws:dataprotection::aws:data-identifier/DriversLicense-US"
        ],
        "Operation" : {
          "Deidentify" : {
            "MaskConfig" : {}
          }
        }
      }
    ]
  })
}

resource "awscc_logs_log_group" "finding" {
  log_group_name = "my-log-group-finding"
}