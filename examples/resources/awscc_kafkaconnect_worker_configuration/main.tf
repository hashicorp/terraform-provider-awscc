# Example Kafka Connect Worker Configuration
resource "awscc_kafkaconnect_worker_configuration" "example" {
  name        = "example-worker-config"
  description = "Example worker configuration for MSK Connect"

  # Base64 encoded content of connect-distributed.properties
  properties_file_content = base64encode(<<-EOT
    key.converter=org.apache.kafka.connect.storage.StringConverter
    value.converter=org.apache.kafka.connect.storage.StringConverter
    EOT
  )

  tags = [{
    key   = "Modified By"
    value = "AWSCC"
  }]
}