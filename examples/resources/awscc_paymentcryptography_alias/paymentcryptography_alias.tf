resource "awscc_paymentcryptography_alias" "example" {
  alias_name = "alias/example"
  key_arn    = awscc_paymentcryptography_key.example.id
}