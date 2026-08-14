resource "awscc_mediaconnect_router_network_interface" "example" {
  name        = "example-router-network-interface"
  region_name = "us-west-2"

  configuration = {
    public = {
      allow_rules = [
        {
          cidr = "10.0.0.0/16"
        },
        {
          cidr = "192.168.1.0/24"
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