# Example MediaLive Network
resource "awscc_medialive_network" "example" {
  name = "example-network"

  ip_pools = [
    {
      cidr = "10.0.0.0/24"
    },
    {
      cidr = "10.0.1.0/24"
    }
  ]

  routes = [
    {
      cidr    = "0.0.0.0/0"
      gateway = "10.0.0.1"
    }
  ]

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