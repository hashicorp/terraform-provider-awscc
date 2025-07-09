resource "awscc_secretsmanager_secret" "example" {
  name        = "example"
  description = "this is a user-provided description of the secret"
}