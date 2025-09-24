resource "awscc_route53resolver_resolver_rule" "example" {
  domain_name          = "example.com"
  name                 = "MyRule"
  resolver_endpoint_id = "rslvr-out-example-id"
  rule_type            = "FORWARD"

  tags = [{
    "key"   = "LineOfBusiness",
    "value" = "Engineering"
  }]

  target_ips = [
    {
      ip   = "192.0.2.6"
      port = "53"
    },
    {
      ip   = "192.0.2.99"
      port = "53"
    }
  ]
}