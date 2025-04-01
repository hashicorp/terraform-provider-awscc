# Create the mail manager traffic policy
resource "awscc_ses_mail_manager_traffic_policy" "example" {
  traffic_policy_name    = "example-traffic-policy"
  default_action         = "ALLOW"
  max_message_size_bytes = 10485760 # 10MB

  policy_statements = [
    {
      action = "DENY"
      conditions = [
        {
          ip_expression = {
            operator = "CIDR_MATCHES"
            values   = ["192.0.2.0/24", "198.51.100.0/24"]
            evaluate = {
              attribute = "SENDER_IP"
            }
          }
        }
      ]
    },
    {
      action = "DENY"
      conditions = [
        {
          string_expression = {
            operator = "EQUALS"
            values   = ["example.com"]
            evaluate = {
              attribute = "RECIPIENT"
            }
          }
        }
      ]
    }
  ]

  tags = [{
    key   = "Modified_By"
    value = "AWSCC"
  }]
}