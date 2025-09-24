# Create a domain list first since it's required for the rules
resource "awscc_route53resolver_firewall_domain_list" "example" {
  name = "example-domain-list"
  domains = [
    "example.com",
    "*.example.org"
  ]
  tags = [{
    key   = "Modified By"
    value = "AWSCC"
  }]
}

# Create the firewall rule group
resource "awscc_route53resolver_firewall_rule_group" "example" {
  name = "example-rule-group"

  firewall_rules = [
    {
      name                    = "example-rule"
      action                  = "ALLOW"
      firewall_domain_list_id = awscc_route53resolver_firewall_domain_list.example.id
      priority                = 100
    }
  ]

  tags = [{
    key   = "Modified By"
    value = "AWSCC"
  }]
}