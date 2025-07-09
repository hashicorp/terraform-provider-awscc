# Create a new DMS migration task
resource "awscc_dms_data_migration_task" "example" {
  replication_task_identifier = "test-dms-task"
  source_endpoint_arn         = "arn:aws:dms:us-west-2:123456789012:endpoint:ABCDEFGHIJKLMNOPQRSTUVWXYZ"
  target_endpoint_arn         = "arn:aws:dms:us-west-2:123456789012:endpoint:ZYXWVUTSRQPONMLKJIHGFEDCBA"
  replication_instance_arn    = "arn:aws:dms:us-west-2:123456789012:rep:ABCDEFGHIJKLMNOPQRSTUVWXYZ"
  migration_type              = "full-load"
  table_mappings = jsonencode({
    rules = [
      {
        rule-type = "selection"
        rule-id   = "1"
        rule-name = "1"
        object-locator = {
          schema-name = "test"
          table-name  = "test"
        }
        rule-action = "include"
      }
    ]
  })

  tags = [{
    key   = "Environment"
    value = "Production"
  }]
}
