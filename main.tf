provider "awscc" {}

resource "awscc_ses_mail_manager_traffic_policy" "main" {
  traffic_policy_name = "test-policy-2"
  default_action      = "DENY"

  policy_statements = [
    {
      action = "ALLOW"
      conditions = [
        {
          ip_expression = {
            operator = "CIDR_MATCHES"
            values   = ["0.0.0.0/0"]
            evaluate = {
              attribute = "SENDER_IP"
            }
          }
        }
      ]
    }
  ]
}

resource "awscc_ses_mail_manager_archive" "main" {
  archive_name = "test-archive-2"

  retention = {
    retention_period = "THREE_MONTHS"
  }
}

resource "awscc_ses_mail_manager_rule_set" "main" {
  rule_set_name = "test-2"
  rules = [
    {
      name = "archive-send"
      actions = [
        {
          archive = {
            target_archive = awscc_ses_mail_manager_archive.main.id
          }
        }
      ]
    }
  ]
}


resource "awscc_ses_mail_manager_ingress_point" "main" {
  ingress_point_name = "test-2"
  traffic_policy_id  = awscc_ses_mail_manager_traffic_policy.main.id
  rule_set_id        = awscc_ses_mail_manager_rule_set.main.id
  type               = "AUTH"
  ingress_point_configuration = {
    smtp_password = "Test12345!"
  }
}