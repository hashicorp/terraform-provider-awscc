resource "awscc_inspectorv2_cis_scan_configuration" "example" {
  scan_name = "example"
  schedule = {
    daily = {
      start_time = {
        time_of_day = "00:00"
        time_zone   = "UTC"
      }
    }
  }

  security_level = "LEVEL_1"
  targets = {
    # use SELF for current account to remove any drift as property transformation is not currently supported.
    account_ids = ["123456789012"]
    target_resource_tags = {
      "Modified By" = ["AWSCC"]
    }

  }

  tags = {
    "Modified By" = "AWSCC"
  }
}
