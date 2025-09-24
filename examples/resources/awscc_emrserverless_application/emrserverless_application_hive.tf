resource "awscc_emrserverless_application" "example" {
  name          = "example"
  release_label = "emr-7.1.0"
  type          = "Hive"

  tags = [{
    key   = "Modified By"
    value = "AWSCC"
  }]
}
