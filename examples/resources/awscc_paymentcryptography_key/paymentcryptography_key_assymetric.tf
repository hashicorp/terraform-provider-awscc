resource "awscc_paymentcryptography_key" "example" {
  exportable = true
  enabled    = true
  key_attributes = {
    key_algorithm = "RSA_2048"
    key_class     = "ASYMMETRIC_KEY_PAIR"
    key_modes_of_use = {
      encrypt = true
      decrypt = true
      wrap    = true
      unwrap  = true
    }
    key_usage = "TR31_D1_ASYMMETRIC_KEY_FOR_DATA_ENCRYPTION"
  }

  tags = [{
    key   = "Modified By"
    value = "AWSCC"
  }]

}