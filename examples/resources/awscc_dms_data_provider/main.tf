
# Create DMS data provider for PostgreSQL
resource "awscc_dms_data_provider" "example" {
  data_provider_name = "example-postgres-provider"
  description        = "Example PostgreSQL Data Provider"
  engine             = "postgres"

  settings = {
    postgre_sql_settings = {
      server_name   = "example-postgres.mydomain.com"
      port          = 5432
      database_name = "example_db"
      ssl_mode      = "require"
    }
  }

  tags = [{
    key   = "Modified By"
    value = "AWSCC"
  }]
}