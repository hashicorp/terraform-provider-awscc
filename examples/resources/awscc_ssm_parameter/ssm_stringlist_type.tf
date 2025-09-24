resource "awscc_ssm_parameter" "example" {
  name  = "commands"
  type  = "StringList"
  value = "date,ls"

  description     = "SSM Parameter of type StringList."
  allowed_pattern = "^[a-zA-Z]{1,10}$"
}