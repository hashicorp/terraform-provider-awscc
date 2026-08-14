resource "awscc_mediaconnect_router_network_interface" "example" {
  name        = "example-router-network-interface"
  region_name = "us-west-2"

  configuration = {
    public = {
      allow_rules = [
        {
          cidr = "10.0.0.0/16"
        }
      ]
    }
  }

  tags = [
    {
      key   = "Name"
      value = "example-router-network-interface"
    },
    {
      key   = "Environment"
      value = "example"
    }
  ]
}

resource "awscc_mediaconnect_router_input" "example" {
  name            = "example-router-input"
  maximum_bitrate = 100000000
  routing_scope   = "REGIONAL"
  tier            = "INPUT_100"

  configuration = {
    standard = {
      protocol              = "RTP"
      network_interface_arn = awscc_mediaconnect_router_network_interface.example.arn
      protocol_configuration = {
        rtp = {
          port                     = 5004
          forward_error_correction = "ENABLED"
        }
      }
    }
  }

  tags = [
    {
      key   = "Environment"
      value = "example"
    },
    {
      key   = "Name"
      value = "example-router-input"
    }
  ]
}
