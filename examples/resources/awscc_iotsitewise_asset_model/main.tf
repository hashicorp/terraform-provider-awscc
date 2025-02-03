# Example of an IoT SiteWise Asset Model for a Wind Turbine
resource "awscc_iotsitewise_asset_model" "wind_turbine" {
  asset_model_name        = "Wind Turbine Model"
  asset_model_description = "Asset model for wind turbines"

  asset_model_properties = [
    # Measurement Properties
    {
      name       = "Wind Speed"
      data_type  = "DOUBLE"
      unit       = "m/s"
      logical_id = "wind_speed_property"
      type = {
        type_name = "Measurement"
      }
    },
    {
      name       = "Power Output"
      data_type  = "DOUBLE"
      unit       = "kW"
      logical_id = "power_output_property"
      type = {
        type_name = "Measurement"
      }
    },
    # Attribute Properties
    {
      name       = "Location"
      data_type  = "STRING"
      logical_id = "location_property"
      type = {
        type_name = "Attribute"
        attribute = {
          default_value = "Unknown"
        }
      }
    }
  ]

  tags = [{
    key   = "Modified By"
    value = "AWSCC"
  }]
}