data "aws_caller_identity" "current" {}
data "aws_region" "current" {}

# Example IoT Domain Configuration
resource "awscc_iot_domain_configuration" "example" {
  domain_configuration_name   = "example-domain-config"
  service_type                = "DATA"
  domain_configuration_status = "DISABLED"

  tags = [{
    key   = "Modified By"
    value = "AWSCC"
  }]
}