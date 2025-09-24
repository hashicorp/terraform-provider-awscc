resource "awscc_rds_option_group" "example_rds_option_group_mssql" {
  engine_name              = "sqlserver-se"
  major_engine_version     = "12.00"
  option_group_description = "SQL Server Option Group"
  tags = [{
    key   = "Modified By"
    value = "AWSCC"
  }]
}