# AWS Ground Station Tracking Config
resource "awscc_groundstation_config" "tracking" {
  name = "example-tracking-config"
  config_data = {
    tracking_config = {
      autotrack = "REQUIRED"
    }
  }
  tags = [
    {
      key   = "Name"
      value = "example-tracking-config"
    }
  ]
}

# AWS Ground Station Antenna Downlink Config
resource "awscc_groundstation_config" "antenna_downlink" {
  name = "example-antenna-downlink-config"
  config_data = {
    antenna_downlink_config = {
      spectrum_config = {
        bandwidth = {
          units = "MHz"
          value = 30
        }
        center_frequency = {
          units = "MHz"
          value = 8200
        }
        polarization = "RIGHT_HAND"
      }
    }
  }
  tags = [
    {
      key   = "Name"
      value = "example-antenna-downlink-config"
    }
  ]
}

# AWS Ground Station Dataflow Endpoint Config
resource "awscc_groundstation_config" "dataflow_endpoint" {
  name = "example-dataflow-endpoint-config"
  config_data = {
    dataflow_endpoint_config = {
      dataflow_endpoint_name = "example-endpoint"
      endpoint = {
        address = {
          name           = "172.31.0.1"
          socket_address = "172.31.0.1"
        }
        mtu  = 1500
        port = 40000
      }
    }
  }
  tags = [
    {
      key   = "Name"
      value = "example-dataflow-endpoint-config"
    }
  ]
}

# Ground Station Mission Profile
resource "awscc_groundstation_mission_profile" "example" {
  name                                    = "example-mission-profile"
  minimum_viable_contact_duration_seconds = 300
  contact_pre_pass_duration_seconds       = 120
  contact_post_pass_duration_seconds      = 180

  tracking_config_arn = awscc_groundstation_config.tracking.arn

  dataflow_edges = [
    {
      source      = awscc_groundstation_config.antenna_downlink.arn
      destination = awscc_groundstation_config.dataflow_endpoint.arn
    }
  ]

  tags = [
    {
      key   = "Environment"
      value = "Development"
    },
    {
      key   = "Name"
      value = "example-mission-profile"
    }
  ]

  depends_on = [
    awscc_groundstation_config.tracking,
    awscc_groundstation_config.antenna_downlink,
    awscc_groundstation_config.dataflow_endpoint
  ]
}
