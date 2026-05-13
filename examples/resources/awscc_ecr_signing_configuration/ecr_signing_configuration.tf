resource "aws_signer_signing_profile" "example" {
  platform_id = "AWSLambda-SHA384-ECDSA"
  name        = "example-signing-profile"

  tags = {
    Name        = "example-signing-profile"
    Environment = "test"
  }
}

resource "awscc_ecr_signing_configuration" "example" {
  signing_profile_parameters = [{
    signing_profile_name = aws_signer_signing_profile.example.name
  }]
  signing_enabled = true
  repository_name = "example-repository"

  tags = [
    {
      key   = "Name"
      value = "example-ecr-signing-config"
    },
    {
      key   = "Environment"
      value = "test"
    }
  ]
}
