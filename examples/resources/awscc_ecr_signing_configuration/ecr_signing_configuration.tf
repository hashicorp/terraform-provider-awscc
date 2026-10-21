resource "aws_signer_signing_profile" "example" {
  platform_id = "Notation-OCI-SHA384-ECDSA"
  name        = "example_ecr_signing_profile"

  tags = {
    Name        = "example_ecr_signing_profile"
    Environment = "test"
  }
}

# The signing configuration applies to the whole registry of the current account and region.
resource "awscc_ecr_signing_configuration" "example" {
  rules = [{
    signing_profile_arn = aws_signer_signing_profile.example.arn
    repository_filters = [{
      filter      = "example-repository*"
      filter_type = "WILDCARD_MATCH"
    }]
  }]
}
