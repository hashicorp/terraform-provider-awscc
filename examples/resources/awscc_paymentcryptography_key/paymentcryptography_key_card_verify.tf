resource "awscc_paymentcryptography_key" "example" {
  exportable = true
  enabled    = true
  key_attributes = {
    key_algorithm = "TDES_2KEY"
    key_class     = "SYMMETRIC_KEY"
    key_modes_of_use = {
      verify   = true
      generate = true
    }
    key_usage = "TR31_C0_CARD_VERIFICATION_KEY"
  }

  tags = [{
    key   = "Modified By"
    value = "AWSCC"
  }]

}