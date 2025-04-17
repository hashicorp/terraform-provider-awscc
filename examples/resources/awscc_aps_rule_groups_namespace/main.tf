# Example rule group data - using a simple recording rule
locals {
  rule_group_data = jsonencode({
    "groups" : [
      {
        "name" : "example_rules",
        "rules" : [
          {
            "record" : "instance:node_cpu_utilization:avg",
            "expr" : "avg by (instance) (rate(node_cpu_seconds_total{mode!=\"idle\"}[5m]) * 100)"
          }
        ]
      }
    ]
  })
}

# Create an AWS Managed Prometheus workspace first
resource "awscc_aps_workspace" "example" {
  alias = "example-workspace"
  tags = [{
    key   = "Modified By"
    value = "AWSCC"
  }]
}

# Create the rule groups namespace
resource "awscc_aps_rule_groups_namespace" "example" {
  name      = "example-rule-groups"
  workspace = awscc_aps_workspace.example.arn
  data      = local.rule_group_data
  tags = [{
    key   = "Modified By"
    value = "AWSCC"
  }]
}