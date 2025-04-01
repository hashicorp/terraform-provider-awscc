# Create the Transfer Profile
resource "awscc_transfer_profile" "example" {
  as_2_id      = "ExampleAs2Id"
  profile_type = "LOCAL"

  # Optional: Add certificates if needed
  # certificate_ids = ["certificate-id-1", "certificate-id-2"]

  tags = [{
    key   = "Modified By"
    value = "AWSCC"
  }]
}