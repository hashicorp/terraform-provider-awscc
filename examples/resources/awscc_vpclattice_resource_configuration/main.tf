resource "awscc_vpclattice_resource_configuration" "example" {
  name                        = "example-resource-configuration"
  port_ranges                 = ["80"]
  protocol_type               = "TCP"
  resource_gateway_id         = awscc_vpclattice_resource_gateway.example.id
  resource_configuration_type = "SINGLE"
  resource_configuration_definition = {
    dns_resource = {
      domain_name     = "example.com"
      ip_address_type = "IPV4"
    }
  }
}