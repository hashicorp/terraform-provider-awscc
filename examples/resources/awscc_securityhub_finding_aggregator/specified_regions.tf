resource "awscc_securityhub_finding_aggregator" "example" {
  region_linking_mode = "SPECIFIED_REGIONS"
  regions = [
    "us-east-1",
    "us-east-2",
    "us-west-1",
    "us-west-2"
  ]
}
