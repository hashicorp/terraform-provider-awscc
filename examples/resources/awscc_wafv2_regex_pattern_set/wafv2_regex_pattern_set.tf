resource "awscc_wafv2_regex_pattern_set" "example" {

  regular_expression_list = [
    "I[a@]mAB[a@]dRequest",
    "^foobar$"
  ]
  name        = "example"
  description = "Example regex pattern set"
  scope       = "REGIONAL"

  tags = [
    {
      key   = "ModifiedBy"
      value = "AWSCC"
    }
  ]
}
