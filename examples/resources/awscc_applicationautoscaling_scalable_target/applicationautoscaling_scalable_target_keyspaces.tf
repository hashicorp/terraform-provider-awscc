resource "awscc_applicationautoscaling_scalable_target" "example" {
  max_capacity       = 20
  min_capacity       = 1
  resource_id        = "keyspace/example/table/${awscc_cassandra_table.example.table_name}"
  scalable_dimension = "cassandra:table:WriteCapacityUnits"
  service_namespace  = "cassandra"
  depends_on         = [awscc_cassandra_table.example]
}


resource "awscc_cassandra_keyspace" "example" {
  keyspace_name = "example"
}


resource "awscc_cassandra_table" "example" {
  table_name    = "example"
  keyspace_name = awscc_cassandra_keyspace.example.id
  partition_key_columns = [{
    column_name = "Message"
    column_type = "ascii"
  }]
  billing_mode = {
    mode = "PROVISIONED"
    provisioned_throughput = {
      read_capacity_units  = 1
      write_capacity_units = 1
    }
  }

}
