# Example of creating a dedicated IP pool for SES
resource "awscc_ses_dedicated_ip_pool" "example" {
  pool_name    = "example-dedicated-ip-pool"
  scaling_mode = "STANDARD"
}