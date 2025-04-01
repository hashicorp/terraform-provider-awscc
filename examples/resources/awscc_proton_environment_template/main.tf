# Main Proton Environment Template resource
resource "awscc_proton_environment_template" "example" {
  name        = "example-environment-template"
  description = "Example environment template created with AWSCC provider"

  display_name = "Example Environment Template"
  provisioning = "CUSTOMER_MANAGED"

  tags = [{
    key   = "Modified By"
    value = "AWSCC"
  }]
}