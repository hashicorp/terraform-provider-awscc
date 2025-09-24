resource "awscc_sso_instance" "example" {
  name = "example"

  tags = [{
    key   = "Modified By"
    value = "AWSCC"
  }]
}