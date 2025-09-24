# Example IVS Playback Restriction Policy
resource "awscc_ivs_playback_restriction_policy" "example" {
  name = "example-restriction-policy"

  # Example: Allow playback only from specific countries (US and Canada)
  allowed_countries = ["US", "CA"]

  # Example: Allow playback only from specific origins
  allowed_origins = [
    "https://example.com",
    "https://test.example.com"
  ]

  # Enable strict origin enforcement
  enable_strict_origin_enforcement = true

  # Tags
  tags = [{
    key   = "Modified By"
    value = "AWSCC"
  }]
}