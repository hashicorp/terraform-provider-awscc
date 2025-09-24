resource "awscc_iot_billing_group" "example" {
  billing_group_name = "example"
  billing_group_properties = {
    billing_group_description = "Example billing group"
  }

  tags = [{
    key   = "Modified By"
    value = "AWSCC"
  }]
}
