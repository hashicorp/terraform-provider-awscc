resource "awscc_omics_sequence_store" "example" {
  name        = "example"
  description = "example"

  tags = {
    "Modified By" = "AWSCC"
  }
}