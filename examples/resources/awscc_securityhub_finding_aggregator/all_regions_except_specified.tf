resource "awscc_securityhub_finding_aggregator" "example" {
  region_linking_mode = "ALL_REGIONS_EXCEPT_SPECIFIED"
  regions = [
    "ap-southeast-1",
    "ap-southeast-2",
    "ap-southeast-3",
    "ap-southeast-4"
  ]
}
