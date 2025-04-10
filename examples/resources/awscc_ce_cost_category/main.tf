# Cost Category for tagging EC2 instances
resource "awscc_ce_cost_category" "example" {
  name          = "EC2-Environment-Categories"
  rule_version  = "CostCategoryExpression.v1"
  default_value = "Other"

  rules = jsonencode([
    {
      Type  = "REGULAR"
      Value = "Development"
      Rule = {
        Tags = {
          Key          = "Environment"
          Values       = ["dev", "development"]
          MatchOptions = ["EQUALS"]
        }
      }
    },
    {
      Type  = "REGULAR"
      Value = "Production"
      Rule = {
        Tags = {
          Key          = "Environment"
          Values       = ["prod", "production"]
          MatchOptions = ["EQUALS"]
        }
      }
    }
  ])

  split_charge_rules = jsonencode([
    {
      Source  = "Production",
      Targets = ["Development"],
      Method  = "PROPORTIONAL"
    }
  ])
}