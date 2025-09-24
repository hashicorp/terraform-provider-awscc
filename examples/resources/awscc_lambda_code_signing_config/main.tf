# Create the Signer signing profile
resource "awscc_signer_signing_profile" "example" {
  platform_id = "AWSLambda-SHA384-ECDSA"
  tags = [{
    key   = "Modified By"
    value = "AWSCC"
  }]
}

# Create the Lambda code signing config
resource "awscc_lambda_code_signing_config" "example" {
  allowed_publishers = {
    signing_profile_version_arns = [awscc_signer_signing_profile.example.profile_version_arn]
  }

  code_signing_policies = {
    untrusted_artifact_on_deployment = "Enforce"
  }

  description = "Example Lambda code signing config using AWSCC"

  tags = [{
    key   = "Modified By"
    value = "AWSCC"
  }]
}