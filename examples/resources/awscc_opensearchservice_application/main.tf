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