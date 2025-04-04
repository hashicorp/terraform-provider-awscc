# Network Firewall Policy
resource "awscc_networkfirewall_firewall_policy" "example" {
  firewall_policy_name = "example-policy"
  description          = "Example Network Firewall Policy"

  firewall_policy = {
    stateless_default_actions = [
      "aws:forward_to_sfe"
    ]
    stateless_fragment_default_actions = [
      "aws:forward_to_sfe"
    ]
    stateful_engine_options = {
      rule_order              = "STRICT_ORDER"
      stream_exception_policy = "DROP"
      flow_timeouts = {
        tcp_idle_timeout_seconds = 1800
      }
    }
    stateful_default_actions = ["aws:drop_strict", "aws:alert_strict"]
  }

  tags = [{
    key   = "Modified By"
    value = "AWSCC"
  }]
}