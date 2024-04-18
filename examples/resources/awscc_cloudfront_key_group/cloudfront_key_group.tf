resource "awscc_cloudfront_key_group" "example" {
  key_group_config = {
    comment = "example key group"
    items   = [awscc_cloudfront_public_key.example.id]
    name    = "example-key-group"
  }
}

resource "awscc_cloudfront_public_key" "example" {
  public_key_config = {
    caller_reference = "test-caller-reference"
    encoded_key      = file("public_key.pem")
    name             = "test_key"
    comment          = "test public key"
  }
}