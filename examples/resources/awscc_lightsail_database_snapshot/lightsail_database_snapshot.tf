resource "aws_lightsail_database" "example" {
  relational_database_name = "example-database"
  availability_zone        = "us-west-2a"
  blueprint_id            = "mysql_8_0"
  bundle_id              = "micro_2_0"
  
  master_database_name = "exampledb"
  master_username      = "exampleuser"
  master_password      = "ExamplePassword123!"  
  
  tags = {
    Environment = "example"
    Name        = "example-database"
  }
}

resource "awscc_lightsail_database_snapshot" "example" {
  relational_database_name          = aws_lightsail_database.example.relational_database_name
  relational_database_snapshot_name = "example-database-snapshot"
  
  tags = [
    {
      key   = "Environment"
      value = "example"
    },
    {
      key   = "Name" 
      value = "example-database-snapshot"
    }
  ]
  
  depends_on = [aws_lightsail_database.example]
}
