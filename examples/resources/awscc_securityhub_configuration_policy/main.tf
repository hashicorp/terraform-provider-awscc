resource "awscc_securityhub_configuration_policy" "example" {
  name        = "SecurityHub-Baseline-Policy"
  description = "Baseline Security Hub configuration policy"

  configuration_policy = {
    security_hub = {
      service_enabled = true
      enabled_standard_identifiers = [
        "aws-foundational-security-best-practices/v/1.0.0",
        "cis-aws-foundations-benchmark/v/1.4.0"
      ]
      security_controls_configuration = {
        enabled_security_control_identifiers = [
          "ACM.1",
          "APIGateway.2",
          "AutoScaling.1",
          "CloudTrail.1",
          "CodeBuild.2"
        ]
        disabled_security_control_identifiers = []
        security_control_custom_parameters = [
          {
            security_control_id = "IAM.1"
            parameters = {
              maxCredentialUsageAge = {
                value = {
                  integer = 90
                }
                value_type = "CUSTOM"
              }
            }
          }
        ]
      }
    }
  }

  tags = [{
    key   = "Modified By"
    value = "AWSCC"
  }]
}