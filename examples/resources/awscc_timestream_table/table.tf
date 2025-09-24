resource "awscc_timestream_table" "this" {
  database_name = "MyTimestreamDB"
  table_name    = "MyTimestreamTable"
  tags = [{
    key   = "Managed By"
    value = "AWSCC"
  }]
}
