resource "awscc_healthlake_fhir_datastore" "example" {
  datastore_name         = "example"
  datastore_type_version = "R4"
  preload_data_config = {
    preload_data_type = "SYNTHEA"
  }
  sse_configuration = {
    kms_encryption_config = {
      cmk_type   = "CUSTOMER_MANAGED_KMS_KEY"
      kms_key_id = awscc_kms_key.example.arn # Use arn of the KMS key here than key_id
    }
  }

  tags = [{
    key   = "Modified By"
    value = "AWSCC"
  }]
}