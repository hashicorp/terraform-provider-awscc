# Get current AWS region and account ID
data "aws_region" "current" {}
data "aws_caller_identity" "current" {}

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

# Output the worker configuration ARN
output "worker_configuration_arn" {
  value = awscc_kafkaconnect_worker_configuration.example.worker_configuration_arn
}