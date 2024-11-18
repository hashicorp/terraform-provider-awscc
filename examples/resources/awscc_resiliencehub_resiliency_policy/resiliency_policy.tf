resource "awscc_resiliencehub_resiliency_policy" "example" {
  policy_name = "example_policy"
  tier        = "MissionCritical"
  policy = {
    software = {
      rpo_in_secs = 900
      rto_in_secs = 3600
    }
    hardware = {
      rpo_in_secs = 300
      rto_in_secs = 300
    }
    region = {
      rpo_in_secs = 300
      rto_in_secs = 300
    }
    az = {
      rpo_in_secs = 300
      rto_in_secs = 300
    }
  }
  policy_description       = "This is an example policy"
  data_location_constraint = "us-west-2"
  tags = [{
    key   = "Modified By"
    value = "AWSCC"
  }]
}





   
