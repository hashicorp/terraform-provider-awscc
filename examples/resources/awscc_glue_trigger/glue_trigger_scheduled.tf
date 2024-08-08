resource "awscc_glue_trigger" "example" {
  name        = "example"
  description = "example for on demand trigger"

  type     = "SCHEDULED"
  schedule = "cron(0 */2 * * ? *)"
  actions = [{
    job_name = "test_job",
    arguments = jsonencode(
      {
        "--job-bookmark-option" : "job-bookmark-enable"
      }
    )
  }]

}