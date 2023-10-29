resource "awscc_route53resolver_resolver_rule_association" "example" {
  resolver_rule_id = awscc_route53resolver_resolver_rule.example.id

  vpc_id = "vpc-example-id"
  name   = "My resolver rule association"
}