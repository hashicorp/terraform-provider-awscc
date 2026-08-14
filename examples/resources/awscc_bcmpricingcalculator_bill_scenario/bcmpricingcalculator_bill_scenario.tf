resource "awscc_bcmpricingcalculator_bill_scenario" "example" {
  name = "example-bill-scenario"

  expires_at = "2027-12-31T23:59:59Z"

  group_sharing_preference = "OPEN"

  tags = [
    {
      key   = "Modified By"
      value = "AWSCC"
    }
  ]
}

output "bill_scenario_id" {
  description = "The ID of the bill scenario"
  value       = awscc_bcmpricingcalculator_bill_scenario.example.id
}
