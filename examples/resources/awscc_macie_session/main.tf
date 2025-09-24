# Create Macie Session
resource "awscc_macie_session" "example" {
  finding_publishing_frequency = "FIFTEEN_MINUTES"
  status                       = "ENABLED"
}