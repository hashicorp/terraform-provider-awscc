resource "awscc_secretsmanager_secret" "example_replica" {
  name = "example_replica"

  replica_regions = [{
    region = "ap-southeast-1"
    },
    {
      region = "ap-southeast-2"
  }]
}