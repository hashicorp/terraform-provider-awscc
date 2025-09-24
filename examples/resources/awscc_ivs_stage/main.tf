# Creates an IVS Stage which enables multiple participants to join a video streaming session 
# and then broadcast that session to the Amazon IVS service for live delivery.
resource "awscc_ivs_stage" "example" {
  name = "example-ivs-stage"

  tags = [{
    key   = "Environment"
    value = "Testing"
    }, {
    key   = "Modified By"
    value = "AWSCC"
  }]
}