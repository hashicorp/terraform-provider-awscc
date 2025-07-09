resource "awscc_controltower_landing_zone" "this" {
  manifest = jsonencode({
    "governedRegions" : toset(["us-west-2", "us-east-1"]),
    "organizationStructure" : {
      "security" : {
        "name" : "Core"
      },
      "sandbox" : {
        "name" : "Sandbox"
      }
    },
    "centralizedLogging" : {
      "accountId" : "YOUR_LOG_ARCHIVE_ACCOUNT_ID",
      "configurations" : {
        "loggingBucket" : {
          "retentionDays" : "60"
        },
        "accessLoggingBucket" : {
          "retentionDays" : "60"
        },
      },
      "enabled" : true
    },
    "securityRoles" : {
      "accountId" : "YOUR_AUDIT_ACCOUNT_ID"
    },
    "accessManagement" : {
      "enabled" : true
    }
    }
  )
  version = "3.3"
}