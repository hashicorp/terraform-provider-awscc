resource "awscc_securityhub_security_control" "example" {
  security_control_id = "ACM.1"
  parameters = {
    daysToExpiration = {
      value = {
        integer = 15
      }
      value_type = "CUSTOM"
    }
  }
  last_update_reason = "Internal compliance requirement"
}
