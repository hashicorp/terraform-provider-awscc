resource "awscc_emrserverless_application" "example" {
  name          = "example"
  release_label = "emr-7.1.0"
  type          = "Spark"
  architecture  = "X86_64"

  network_configuration = {
    security_group_ids = ["sg-0f75e03716b3442c0"]
    subnet_ids         = ["subnet-0e3cd1df31dgb5e6c"] # private subnet
  }
  auto_start_configuration = {
    enabled = true
  }
  auto_stop_configuration = {
    enabled              = true
    idle_timeout_minutes = 10
  }

  tags = [{
    key   = "Modified By"
    value = "AWSCC"
  }]
}
