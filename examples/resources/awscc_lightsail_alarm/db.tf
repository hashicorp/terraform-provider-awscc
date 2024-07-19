resource "awscc_lightsail_database" "example" {
  master_database_name             = "example"
  master_username                  = "admin"
  relational_database_blueprint_id = "mysql_8_0"
  relational_database_bundle_id    = "micro_2_0"
  relational_database_name         = "example-db"
}

resource "awscc_lightsail_alarm" "example" {
  alarm_name              = "example-alarm"
  comparison_operator     = "LessThanThreshold"
  evaluation_periods      = 1
  metric_name             = "FreeStorageSpace"
  monitored_resource_name = awscc_lightsail_database.example.relational_database_name
  threshold               = 10737418240
  notification_triggers   = ["ALARM"]
}
