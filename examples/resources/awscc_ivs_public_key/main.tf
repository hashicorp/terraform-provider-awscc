# Data sources for dynamic values
data "aws_caller_identity" "current" {}
data "aws_region" "current" {}

# Create IVS Public Key resource
resource "awscc_ivs_public_key" "example" {
  name = "example-ivs-public-key"
  # This is an example RSA public key, you should replace it with your actual public key
  public_key_material = "-----BEGIN PUBLIC KEY-----\nMIIBIjANBgkqhkiG9w0BAQEFAAOCAQ8AMIIBCgKCAQEAjeMQF6KuSCiiWF3Owc5C\nKq3DC3hSIgdaeBUAL5qQvRLaQ4/XEktOzucM64ueUxE8Fa6wITEWKHLT2B1Tc0Ni\nrCcATZqJB5xVcB5AMyGLb5H6HrVuPRiuf9ewXHbk+8FvhPe9cjWki5QV7ERm0Z6z\nM4RXBvhECRsxYt9bluyfod6MRQRlST/L13pkB6mYhxqZWA2t+r+04hdK6EP20MvG\nSVVzXKD2+Gtg7ZVBlH5bzU7pQc6w5jJr0hppAHY8gnHR31twhH92qpAIHjSYPfHg\nJqXzYYHlR5XQPvmEXbyHKryF2G0E8Su0XQqGOBa0bWpjEje1f9tD/vkAEE1jnR47\nKwIDAQAB\n-----END PUBLIC KEY-----"
  tags = [{
    key   = "Modified By"
    value = "AWSCC"
  }]
}