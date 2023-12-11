resource "awscc_ssm_parameter" "example" {
  name            = "command"
  type            = "String"
  value           = "date"
  description     = "SSM Parameter for running date command."
  allowed_pattern = "^[a-zA-Z]{1,10}$"
}