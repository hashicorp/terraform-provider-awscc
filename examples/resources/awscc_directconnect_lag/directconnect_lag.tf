resource "awscc_directconnect_lag" "example" {
  lag_name              = "example-lag"
  location              = "EqSe2-EQ"
  connections_bandwidth = "10Gbps"

  tags = [
    {
      key   = "Environment"
      value = "example"
    },
    {
      key   = "Name"
      value = "example-lag"
    }
  ]
}
