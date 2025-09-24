# First create a signal catalog as it's required for state template
resource "awscc_iotfleetwise_signal_catalog" "example" {
  name        = "example-signal-catalog"
  description = "Example signal catalog for state template"
  nodes = [
    {
      actuator  = null
      attribute = null
      branch = {
        fully_qualified_name = "Vehicle"
      }
      description = "Vehicle branch"
      sensor      = null
    },
    {
      actuator  = null
      attribute = null
      branch = {
        fully_qualified_name = "Vehicle.Speed"
      }
      description = "Vehicle speed information"
      sensor      = null
    },
    {
      actuator    = null
      attribute   = null
      branch      = null
      description = "Vehicle speed in km/h"
      sensor = {
        data_type            = "DOUBLE"
        fully_qualified_name = "Vehicle.Speed.Chassis"
        max                  = 300
        min                  = 0
        unit                 = "km/h"
      }
    }
  ]
  tags = [{
    key   = "Modified By"
    value = "AWSCC"
  }]
}

# Create the state template
resource "awscc_iotfleetwise_state_template" "example" {
  name               = "example-state-template"
  description        = "Example IoT FleetWise state template"
  signal_catalog_arn = awscc_iotfleetwise_signal_catalog.example.arn
  state_template_properties = [
    "Vehicle.Speed.Chassis"
  ]
  data_extra_dimensions = [
    "Vehicle"
  ]
  metadata_extra_dimensions = [
    "Vehicle.Speed"
  ]
  tags = [{
    key   = "Modified By"
    value = "AWSCC"
  }]
}