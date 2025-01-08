# IoT FleetWise Signal Catalog example
resource "awscc_iotfleetwise_signal_catalog" "example" {
  name        = "example-signal-catalog"
  description = "Example signal catalog for IoT FleetWise"

  # Example nodes with different types
  nodes = [
    # Branch node example
    {
      branch = {
        description          = "Vehicle Data"
        fully_qualified_name = "Vehicle"
      }
    },
    # Sensor node example
    {
      sensor = {
        fully_qualified_name = "Vehicle.Speed"
        description          = "Vehicle speed sensor"
        data_type            = "DOUBLE"
        unit                 = "km/h"
        min                  = 0
        max                  = 300
      }
    },
    # Actuator node example
    {
      actuator = {
        fully_qualified_name = "Vehicle.Brake"
        description          = "Vehicle brake status"
        data_type            = "BOOLEAN"
        allowed_values       = ["true", "false"]
      }
    },
    # Attribute node example
    {
      attribute = {
        fully_qualified_name = "Vehicle.VIN"
        description          = "Vehicle Identification Number"
        data_type            = "STRING"
      }
    }
  ]

  tags = [{
    key   = "Modified By"
    value = "AWSCC"
  }]
}