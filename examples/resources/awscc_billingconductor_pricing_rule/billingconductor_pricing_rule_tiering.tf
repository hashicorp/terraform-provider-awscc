resource "awscc_billingconductor_pricing_rule" "example_tiering" {
  name        = "Test-Tiering"
  description = "Setup tiering."
  scope       = "GLOBAL" #"GLOBAL" "SERVICE" "BILLING_ENTITY" "SKU"]
  type        = "TIERING"

  tiering = {
    free_tier = {
      activated = false
    }
  }


  tags = [
    {
      key   = "Modified By"
      value = "AWSCC"
    }

  ]
}