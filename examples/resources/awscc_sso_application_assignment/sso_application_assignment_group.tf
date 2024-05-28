resource "awscc_sso_application_assignment" "example2" {
  application_arn = awscc_sso_application.example.application_arn
  principal_id    = var.group_id
  principal_type  = "GROUP"
}

variable "group_id" {
  type = string
}