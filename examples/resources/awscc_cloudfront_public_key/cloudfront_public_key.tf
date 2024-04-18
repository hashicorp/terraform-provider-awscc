resource "awscc_cloudfront_public_key" "example" {
  public_key_config = {
    caller_reference = "test-caller-reference"
    encoded_key      = file("public_key.pem")
    name             = "test_key"
    comment          = "test public key"
  }
}