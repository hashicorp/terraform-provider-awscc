# Get current AWS region and account ID
data "aws_region" "current" {}
data "aws_caller_identity" "current" {}

# Example service principal name with both connector ARN and directory registration ARN
resource "awscc_pcaconnectorad_service_principal_name" "this" {
  connector_arn              = "arn:aws:pca-connector-ad:${data.aws_region.current.name}:${data.aws_caller_identity.current.account_id}:connector/12345678-1234-1234-1234-123456789012"
  directory_registration_arn = "arn:aws:pca-connector-ad:${data.aws_region.current.name}:${data.aws_caller_identity.current.account_id}:directory-registration/d-0123456789"
}