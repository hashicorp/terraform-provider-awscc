data "aws_caller_identity" "current" {}
data "aws_region" "current" {}

# Create MediaLive SDI Source
resource "awscc_medialive_sdi_source" "example" {
  name = "example-sdi-source"
  type = "SINGLE"

  tags = [{
    key   = "Modified By"
    value = "AWSCC"
  }]
}