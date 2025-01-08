# MSK Configuration
resource "awscc_msk_configuration" "example" {
  name = "example-msk-config"

  # Server properties in the format of a string of key-value pairs
  # These are common Apache Kafka configuration parameters
  server_properties = <<EOT
auto.create.topics.enable=true
default.replication.factor=3
min.insync.replicas=2
num.partitions=1
log.retention.hours=168
EOT

  description         = "Example MSK configuration"
  kafka_versions_list = ["2.8.1", "3.3.1"]
}