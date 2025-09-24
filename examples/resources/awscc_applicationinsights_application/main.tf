# Create resource group first
resource "awscc_resourcegroups_group" "example" {
  name = "my-ai-example-rg"

  resource_query = {
    type = "TAG_FILTERS_1_0"
    query = {
      resource_type_filters = ["AWS::EC2::Instance"]
      tag_filters = [
        {
          key    = "Stage"
          values = ["Production"]
        }
      ]
    }
  }

  tags = [{
    key   = "Modified By"
    value = "AWSCC"
  }]
}

# Create Application Insights application
resource "awscc_applicationinsights_application" "example" {
  resource_group_name = awscc_resourcegroups_group.example.name

  auto_configuration_enabled = true
  cwe_monitor_enabled        = true
  ops_center_enabled         = true

  tags = [{
    key   = "Modified By"
    value = "AWSCC"
  }]

  # Optional: Custom log pattern set
  log_pattern_sets = [{
    pattern_set_name = "ExamplePatternSet"
    log_patterns = [{
      pattern_name = "ErrorPattern"
      pattern      = "Error.*"
      rank         = 1
    }]
  }]
}