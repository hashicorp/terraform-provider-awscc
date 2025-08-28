# Create a Traffic Mirror Filter
resource "awscc_ec2_traffic_mirror_filter" "example" {
  description = "Example Traffic Mirror Filter"
  tags = [
    {
      key   = "Name"
      value = "example-filter"
    },
    {
      key   = "Environment"
      value = "test"
    }
  ]
}

# Create Traffic Mirror Filter Rule
resource "awscc_ec2_traffic_mirror_filter_rule" "example" {
  description              = "Example Traffic Mirror Filter Rule"
  traffic_mirror_filter_id = awscc_ec2_traffic_mirror_filter.example.id

  # Rule Configuration
  traffic_direction = "ingress"
  rule_number       = 100
  rule_action       = "accept"

  # Network Configuration
  source_cidr_block      = "10.0.0.0/16"
  destination_cidr_block = "10.0.0.0/16"
  protocol               = 17 # UDP protocol

  # Optional port ranges
  source_port_range = {
    from_port = 10
    to_port   = 50
  }

  destination_port_range = {
    from_port = 50
    to_port   = 100
  }

  # Tags
  tags = [
    {
      key   = "Name"
      value = "example-filter-rule"
    },
    {
      key   = "Environment"
      value = "test"
    }
  ]
}
