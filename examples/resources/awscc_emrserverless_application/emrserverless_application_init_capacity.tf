resource "awscc_emrserverless_application" "example" {
  name          = "example"
  release_label = "emr-7.1.0"
  type          = "Hive"
  initial_capacity = [
    {
      key = "HiveDriver"
      value = {
        worker_count = 1
        worker_configuration = {
          cpu    = "2 vCPU"
          memory = "10 GB"
        }
      }
    }
  ]

  tags = [{
    key   = "Modified By"
    value = "AWSCC"
  }]
}
