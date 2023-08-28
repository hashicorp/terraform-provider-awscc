resource "awscc_cassandra_keyspace" "awscc_cassandra_keyspace_example" {
  keyspace_name = "awscc_cassandra_keyspace_example"
}

resource "awscc_cassandra_table" "awscc_cassandra_table_example" {
  keyspace_name = "awscc_cassandra_keyspace_example"
  partition_key_columns = [{
    column_name = "Message"
    column_type = "ASCII"
  }]
  depends_on = [awscc_cassandra_keyspace.awscc_cassandra_keyspace_example]
}