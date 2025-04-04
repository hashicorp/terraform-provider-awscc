# Create the Evidently Project first
resource "awscc_evidently_project" "example" {
  name        = "example-project"
  description = "Example project for launch"

  tags = [{
    key   = "Modified By"
    value = "AWSCC"
  }]
}

# Create an Evidently Feature (required for the launch)
resource "awscc_evidently_feature" "example" {
  name        = "example-feature"
  project     = awscc_evidently_project.example.name
  description = "Example feature for launch"

  variations = [{
    variation_name = "Variation1"
    string_value   = jsonencode({ "flag" = true })
    description    = "First variation"
    },
    {
      variation_name = "Variation2"
      string_value   = jsonencode({ "flag" = false })
      description    = "Second variation"
  }]

  default_variation = "Variation1"

  tags = [{
    key   = "Modified By"
    value = "AWSCC"
  }]
}

# Create the Evidently Launch
resource "awscc_evidently_launch" "example" {
  name        = "example-launch"
  project     = awscc_evidently_project.example.name
  description = "Example launch configuration"

  groups = [
    {
      feature     = awscc_evidently_feature.example.name
      group_name  = "Control"
      variation   = "Variation1"
      description = "Control group"
    },
    {
      feature     = awscc_evidently_feature.example.name
      group_name  = "Treatment"
      variation   = "Variation2"
      description = "Treatment group"
    }
  ]

  scheduled_splits_config = [
    {
      start_time = "2025-01-03T00:00:00Z"
      group_weights = [
        {
          group_name   = "Control"
          split_weight = 50
        },
        {
          group_name   = "Treatment"
          split_weight = 50
        }
      ]
    }
  ]

  execution_status = {
    status        = "START"
    desired_state = "COMPLETED"
  }

  metric_monitors = [
    {
      metric_name = "ClickCount"
      event_pattern = jsonencode({
        "eventType" : ["click"]
      })
      unit_label    = "Clicks"
      value_key     = "$.value"
      entity_id_key = "$.user_id"
    }
  ]

  tags = [{
    key   = "Modified By"
    value = "AWSCC"
  }]
}