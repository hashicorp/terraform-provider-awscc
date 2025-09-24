# Enable IoT Fleet Indexing
resource "aws_iot_indexing_configuration" "example" {
  thing_indexing_configuration {
    thing_connectivity_indexing_mode = "STATUS"
    thing_indexing_mode              = "REGISTRY_AND_SHADOW"
  }
}

# Add a time delay for indexing to complete
resource "time_sleep" "wait_60_seconds" {
  depends_on      = [aws_iot_indexing_configuration.example]
  create_duration = "60s"
}

# Create an IoT Fleet Metric
resource "awscc_iot_fleet_metric" "example" {
  depends_on        = [time_sleep.wait_60_seconds]
  metric_name       = "device-connectivity-metric"
  description       = "Fleet metric to track device connectivity status"
  index_name        = "AWS_Things"
  query_string      = "connectivity.connected:true"
  query_version     = "2017-09-30"
  aggregation_field = "connectivity.connected"
  period            = 300 # 5 minutes

  aggregation_type = {
    name = "Statistics"
    values = [
      "count"
    ]
  }

  unit = "Count"

  tags = [{
    key   = "Modified By"
    value = "AWSCC"
  }]
}