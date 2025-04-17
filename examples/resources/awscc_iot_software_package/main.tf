# IoT Software Package resource
resource "awscc_iot_software_package" "example" {
  package_name = "example-package"
  description  = "Example IoT software package created using AWSCC provider"

  tags = [{
    key   = "Modified By"
    value = "AWSCC"
  }]
}