resource "awscc_sqs_queue" "terraform_awscc_queue_kms" {
  queue_name                        = "terraform-awscc-queue-kms-example"
  kms_master_key_id                 = "alias/aws/sqs"
  kms_data_key_reuse_period_seconds = 300
  tags = [{
    key   = "Modified By"
    value = "AWSCC"
  }]

}
