data "aws_caller_identity" "current" {}
data "aws_region" "current" {}

# Example OpenSearch Service application
resource "awscc_opensearchservice_application" "example" {
  name = "exampleapp123"

  tags = [
    {
      key   = "Modified By"
      value = "AWSCC"
    }
  ]
}