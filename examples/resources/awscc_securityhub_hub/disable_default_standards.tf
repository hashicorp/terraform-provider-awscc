resource "awscc_securityhub_hub" "example" {
  auto_enable_controls      = true
  control_finding_generator = "SECURITY_CONTROL"
  enable_default_standards  = false
}
