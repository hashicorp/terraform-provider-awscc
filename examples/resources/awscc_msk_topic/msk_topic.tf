resource "awscc_msk_topic" "example" {
  cluster_arn        = "arn:aws:kafka:us-west-2:204034886740:cluster/example-msk-cluster/fd3f8bae-c922-47b4-90e6-eecf2e212a03-2"
  topic_name         = "example-topic"
  partition_count    = 3
  replication_factor = 2
}
