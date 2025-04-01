# Create a service template
resource "awscc_proton_service_template" "example" {
  name         = "example-service-template"
  display_name = "Example Service Template"
  description  = "Example service template created with AWSCC provider"

  # Set pipeline provisioning to CUSTOMER_MANAGED
  pipeline_provisioning = "CUSTOMER_MANAGED"

  # Add tags
  tags = [{
    key   = "Modified By"
    value = "AWSCC"
  }]
}