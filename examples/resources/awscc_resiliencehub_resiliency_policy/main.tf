resource "awscc_resiliencehub_resiliency_policy" "example" {
  policy_name        = "example-policy"
  policy_description = "Example Resiliency Policy"
  tier               = "MissionCritical"

  policy = {
    az = {
      rpo_in_secs = 300 # 5 minutes
      rto_in_secs = 600 # 10 minutes
    }
    hardware = {
      rpo_in_secs = 600  # 10 minutes
      rto_in_secs = 1200 # 20 minutes
    }
    software = {
      rpo_in_secs = 300 # 5 minutes
      rto_in_secs = 600 # 10 minutes
    }
    region = {
      rpo_in_secs = 3600 # 1 hour
      rto_in_secs = 7200 # 2 hours
    }
  }

  data_location_constraint = "AnyLocation"

  tags = [{
    key   = "Modified By"
    value = "AWSCC"
  }, {
    key   = "Environment"
    value = "example"
  }]
}
