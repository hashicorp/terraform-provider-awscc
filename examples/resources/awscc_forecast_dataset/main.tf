# Example Forecast Dataset
resource "awscc_forecast_dataset" "example" {
  dataset_name   = "exampleForecastDataset"
  dataset_type   = "TARGET_TIME_SERIES"
  domain         = "RETAIL"
  data_frequency = "D" # Daily frequency

  schema = {
    attributes = [
      {
        attribute_name = "timestamp"
        attribute_type = "timestamp"
      },
      {
        attribute_name = "demand"
        attribute_type = "float"
      },
      {
        attribute_name = "item_id"
        attribute_type = "string"
      }
    ]
  }

  tags = [
    {
      key   = "Environment"
      value = "Production"
    },
    {
      key   = "Modified By"
      value = "AWSCC"
    }
  ]
}