resource "awscc_refactorspaces_environment" "example" {
  name                = "example-environment"
  description         = "Example RefactorSpaces Environment created with AWSCC"
  network_fabric_type = "NONE"

  tags = [{
    key   = "Modified By"
    value = "AWSCC"
  }]
}