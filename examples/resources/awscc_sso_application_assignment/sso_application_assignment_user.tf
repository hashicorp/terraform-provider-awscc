resource "awscc_sso_application_assignment" "example" {
  application_arn = awscc_sso_application.example.application_arn
  principal_id    = var.user_id
  principal_type  = "USER"
}

variable "user_id" {
  type = string
}