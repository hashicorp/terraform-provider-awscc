resource "awscc_signer_signing_profile" "example" {
  platform_id = "AWSLambda-SHA384-ECDSA"

  signature_validity_period = {
    value = 5
    type  = "YEARS"
  }

}