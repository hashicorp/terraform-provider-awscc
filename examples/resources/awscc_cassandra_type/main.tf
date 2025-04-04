# Create a keyspace first
resource "awscc_cassandra_keyspace" "example" {
  keyspace_name = "example_keyspace"
}

# Example Cassandra Type
resource "awscc_cassandra_type" "example" {
  keyspace_name = awscc_cassandra_keyspace.example.keyspace_name
  type_name     = "address_type"

  fields = [{
    field_name = "street"
    field_type = "text"
    }, {
    field_name = "city"
    field_type = "text"
    }, {
    field_name = "postal_code"
    field_type = "text"
    }, {
    field_name = "country"
    field_type = "text"
  }]
}