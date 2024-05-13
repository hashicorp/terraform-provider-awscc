resource "awscc_signer_signing_profile" "example" {
  platform_id = "AWSLambda-SHA384-ECDSA"
}

resource "awscc_signer_profile_permission" "example" {
  profile_name = awscc_signer_signing_profile.example.profile_name
  action       = "signer:GetSigningProfile"
  principal    = var.role_arn
  statement_id = "statement_example"
}

variable "role_arn" {
  type = string
}