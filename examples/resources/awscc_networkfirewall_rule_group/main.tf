resource "awscc_networkfirewall_rule_group" "example" {
  capacity        = 100
  rule_group_name = "example-stateful-rule-group"
  type            = "STATEFUL"

  rule_group = {
    rule_variables = {
      ip_sets = {
        HOME_NET = {
          definition = ["10.0.0.0/16", "192.168.0.0/16"]
        }
      }
      port_sets = {
        HTTP_PORTS = {
          definition = ["80", "8080"]
        }
      }
    }
    rules_source = {
      stateful_rules = [{
        action = "DROP"
        header = {
          destination      = "ANY"
          destination_port = "ANY"
          direction        = "ANY"
          protocol         = "TCP"
          source           = "ANY"
          source_port      = "ANY"
        }
        rule_options = [{
          keyword  = "sid"
          settings = ["1"]
        }]
      }]
    }
  }

  tags = [{
    key   = "Modified By"
    value = "AWSCC"
  }]
}