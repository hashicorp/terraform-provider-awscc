resource "awscc_billingconductor_pricing_rule" "example_tiering" {
  name        = "EnableFreeTiering"
  
  scope       = "GLOBAL"
  type        = "TIERING"

  tiering = {
    free_tier = {
      activated = true
    }
  }


  tags = [
    {
      key   = "Modified By"
      value = "AWSCC"
    }

  ]
}