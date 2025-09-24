# First, let's create an asset model
resource "awscc_iotsitewise_asset_model" "example" {
  asset_model_name        = "Example Asset Model"
  asset_model_description = "An example asset model for documentation"

  asset_model_properties = [{
    data_type  = "STRING"
    logical_id = "ModelProperty1"
    name       = "Property1"
    type = {
      attribute = {
        default_value = "default"
      }
      type_name = "Attribute"
    }
  }]

  tags = [{
    key   = "Modified By"
    value = "AWSCC"
  }]
}

# Then create the asset using the model
resource "awscc_iotsitewise_asset" "example" {
  asset_name        = "Example Asset"
  asset_model_id    = awscc_iotsitewise_asset_model.example.asset_model_id
  asset_description = "An example IoT SiteWise asset"

  asset_properties = [{
    logical_id         = "ModelProperty1"
    notification_state = "ENABLED"
  }]

  tags = [{
    key   = "Modified By"
    value = "AWSCC"
  }]
}