resource "awscc_paymentcryptography_key" "example" {
  exportable = true
  enabled    = true
  key_attributes = {
    key_algorithm = "TDES_3KEY"
    key_class     = "SYMMETRIC_KEY"
    key_modes_of_use = {
      generate = true
      verify   = true
    }
    key_usage = "TR31_V2_VISA_PIN_VERIFICATION_KEY"
  }

  tags = [{
    key   = "Modified By"
    value = "AWSCC"
  }]

}