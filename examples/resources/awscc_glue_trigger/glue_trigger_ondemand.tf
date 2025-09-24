resource "awscc_glue_trigger" "example" {
  name        = "example"
  description = "example for on demand trigger"

  type = "ON_DEMAND"
  actions = [{
    job_name = "test_job"
  }]

}
