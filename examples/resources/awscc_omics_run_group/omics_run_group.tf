resource "awscc_omics_run_group" "example" {
  name         = "example-run-group"
  max_cpus     = 100
  max_duration = 1440 # 24 hours in minutes
  max_runs     = 10
  max_gpus     = 2

  tags = {
    Environment = "example"
    Name        = "example-run-group"
  }
}
