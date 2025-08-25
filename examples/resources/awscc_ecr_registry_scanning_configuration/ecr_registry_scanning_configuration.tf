# ECR Registry Scanning Configuration resource
resource "awscc_ecr_registry_scanning_configuration" "example" {
  # Configure Enhanced scanning type
  scan_type = "ENHANCED"

  # Configure scanning rules
  rules = [
    {
      scan_frequency = "CONTINUOUS_SCAN"
      repository_filters = [
        {
          filter      = "*"
          filter_type = "WILDCARD"
        }
      ]
    }
  ]
}
