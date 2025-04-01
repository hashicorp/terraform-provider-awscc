# Example SLO for API Latency
resource "awscc_applicationsignals_service_level_objective" "example" {
  name        = "example-api-latency-slo"
  description = "API latency SLO example using CloudWatch metrics"

  # Set up a period-based SLI
  sli = {
    comparison_operator = "LessThan"
    metric_threshold    = 1000 # 1000ms latency threshold

    sli_metric = {

      metric_data_queries = [{
        id = "latency"
        metric_stat = {
          metric = {
            namespace   = "AWS/ApiGateway"
            metric_name = "Latency"
            dimensions = [
              {
                name  = "ApiName"
                value = "ExampleAPI"
              }
            ]
          }
          period = 300
          stat   = "p90"
        }
      }]
    }
  }

  # Define the goal
  goal = {
    attainment_goal = 99.9
    interval = {
      rolling_interval = {
        duration      = 7
        duration_unit = "DAY"
      }
    }
  }

  # Add burn rate configurations
  burn_rate_configurations = [
    {
      look_back_window_minutes = 60
    },
    {
      look_back_window_minutes = 1440 # 24 hours
    }
  ]

  # Add tags
  tags = [
    {
      key   = "Environment"
      value = "Production"
    },
    {
      key   = "ModifiedBy"
      value = "AWSCC"
    }
  ]
}