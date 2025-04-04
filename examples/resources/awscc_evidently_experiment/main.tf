# Create an Evidently Project first since it's required for the experiment
resource "awscc_evidently_project" "example" {
  name        = "example-project"
  description = "Example project for Evidently experiment"
  tags = [{
    key   = "ModifiedBy"
    value = "AWSCC"
  }]
}

# Create the Evidently Experiment
resource "awscc_evidently_experiment" "example" {
  name        = "example-experiment"
  description = "Example experiment using AWSCC provider"
  project     = awscc_evidently_project.example.name

  metric_goals = [
    {
      desired_change = "INCREASE"
      entity_id_key  = "userDetails.userId"
      metric_name    = "page_load_time"
      value_key      = "details.timeInMs"
      event_pattern = jsonencode({
        detail-type : ["page-load"]
        source : ["com.example.web"]
      })
      unit_label = "milliseconds"
    }
  ]

  online_ab_config = {
    control_treatment_name = "control"
    treatment_weights = [
      {
        treatment    = "treatment-A"
        split_weight = 40
      },
      {
        treatment    = "control"
        split_weight = 60
      }
    ]
  }

  treatments = [
    {
      feature        = "page-rendering"
      treatment_name = "control"
      description    = "Current page rendering logic"
      variation      = "original"
    },
    {
      feature        = "page-rendering"
      treatment_name = "treatment-A"
      description    = "New optimized page rendering"
      variation      = "optimized"
    }
  ]

  sampling_rate = 10

  running_status = {
    status = "START"
  }

  tags = [{
    key   = "ModifiedBy"
    value = "AWSCC"
  }]
}