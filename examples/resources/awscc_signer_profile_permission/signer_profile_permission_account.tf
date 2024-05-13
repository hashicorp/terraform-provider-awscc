resource "awscc_signer_signing_profile" "example" {
  platform_id = "AWSLambda-SHA384-ECDSA"
}

resource "awscc_signer_profile_permission" "example" {
  profile_name = awscc_signer_signing_profile.example.profile_name
  action       = "signer:StartSigningJob"
  principal    = var.account_id
  statement_id = "statement_example"
}

variable "account_id" {
  type = string
}