# Example IVS Channel resource
resource "awscc_ivs_channel" "example" {
  name            = "demo-channel"
  type            = "STANDARD"
  latency_mode    = "LOW"
  authorized      = true
  insecure_ingest = false

  tags = [{
    key   = "Modified By"
    value = "AWSCC"
  }]
}