# First create a signal catalog which is required for model manifest
resource "awscc_iotfleetwise_signal_catalog" "example" {
  name        = "example-signal-catalog"
  description = "Example signal catalog for model manifest"
  nodes = [
    {
      branch = {
        fully_qualified_name = "Vehicle"
        description          = "Vehicle branch"
      }
    },
    {
      sensor = {
        fully_qualified_name = "Vehicle.Speed"
        data_type            = "DOUBLE"
        description          = "Vehicle Speed in km/h"
        unit                 = "km/h"
      }
    }
  ]

  tags = [{
    key   = "Modified By"
    value = "AWSCC"
  }]
}

# Create the model manifest
resource "awscc_iotfleetwise_model_manifest" "example" {
  name               = "example-model-manifest"
  description        = "Example model manifest"
  signal_catalog_arn = awscc_iotfleetwise_signal_catalog.example.arn
  nodes = [
    "Vehicle.Speed"
  ]
  status = "ACTIVE"

  tags = [{
    key   = "Modified By"
    value = "AWSCC"
  }]
}