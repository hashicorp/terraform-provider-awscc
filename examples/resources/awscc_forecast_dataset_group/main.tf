# Example Forecast Dataset Group
resource "awscc_forecast_dataset_group" "example" {
  dataset_group_name = "exampleForecastDatasetGroup"
  domain             = "RETAIL" # Valid domains: RETAIL | CUSTOM | INVENTORY_PLANNING | EC2_CAPACITY | WORK_FORCE | WEB_TRAFFIC | METRICS
  tags = [{
    key   = "Modified By"
    value = "AWSCC"
  }]
}