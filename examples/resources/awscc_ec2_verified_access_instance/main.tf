# Define a log group for CloudWatch Logs
resource "awscc_logs_log_group" "vac_logs" {
  log_group_name    = "/aws/verified-access/instance-logs"
  retention_in_days = 7

  tags = [{
    key   = "Modified By"
    value = "AWSCC"
  }]
}

# Create the Verified Access Instance
resource "awscc_ec2_verified_access_instance" "example" {
  description  = "Example Verified Access Instance with logging"
  fips_enabled = true

  logging_configurations = {
    cloudwatch_logs = {
      enabled   = true
      log_group = awscc_logs_log_group.vac_logs.log_group_name
    }
    include_trust_context = true
    log_version           = "ocsf-1.0.0-rc.2"
  }

  tags = [{
    key   = "Modified By"
    value = "AWSCC"
  }]
}