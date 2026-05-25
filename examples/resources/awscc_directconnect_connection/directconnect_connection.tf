resource "awscc_directconnect_connection" "example" {
  bandwidth       = "1Gbps"
  connection_name = "example-connection"
  location        = "EqSe2-EQ"

  tags = [
    {
      key   = "Name"
      value = "example-connection"
    },
    {
      key   = "Environment"
      value = "example"
    }
  ]
}
