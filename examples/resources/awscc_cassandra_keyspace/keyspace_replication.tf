resource "awscc_cassandra_keyspace" "awscc_cassandra_example" {
  keyspace_name = "awscc_cassandra_example"
  replication_specification = {
    replication_strategy = "MULTI_REGION"
    region_list          = ["us-west-2", "us-east-1"]
  }
}