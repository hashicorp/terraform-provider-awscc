resource "awscc_cassandra_keyspace" "awscc_cassandra_example" {
  keyspace_name = "awscc_cassandra_example"
  tags = [{
    "key"   = "Name"
    "value" = "awcc_example"
    },
    {
      "key"   = "Type"
      "value" = "Casandra"
  }]
}