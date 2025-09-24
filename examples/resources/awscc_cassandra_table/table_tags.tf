resource "awscc_cassandra_keyspace" "awscc_cassandra_keyspace_example" {
  keyspace_name = "awscc_cassandra_keyspace_example"
}

resource "awscc_cassandra_table" "awscc_cassandra_table_example" {
  keyspace_name = awscc_cassandra_keyspace.awscc_cassandra_keyspace_example.id
  table_name    = "awscc_cassandra_table_example"
  partition_key_columns = [{
    column_name = "Message"
    column_type = "ascii"
  }]
  regular_columns = [{
    column_name = "name"
    column_type = "TEXT"
    }, {
    column_name = "region"
    column_type = "TEXT"
    }, {
    column_name = "role"
    column_type = "TEXT"
    }, {
    column_name = "vacation_hrs"
    column_type = "FLOAT"
    }
  ]
  tags = [{
    "key"   = "Name"
    "value" = "awscc"
    },
    {
      "key"   = "Components"
      "value" = "Cassandra"
  }]
}