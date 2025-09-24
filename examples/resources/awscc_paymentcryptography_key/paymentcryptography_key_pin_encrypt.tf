resource "awscc_paymentcryptography_key" "example" {
  exportable = true
  enabled    = true
  key_attributes = {
    key_algorithm = "TDES_3KEY"
    key_class     = "SYMMETRIC_KEY"
    key_modes_of_use = {
      encrypt = true
      decrypt = true
      wrap    = true
      unwrap  = true
    }
    key_usage = "TR31_P0_PIN_ENCRYPTION_KEY"
  }

  tags = [{
    key   = "Modified By"
    value = "AWSCC"
  }]

}