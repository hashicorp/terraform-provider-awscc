# Get current AWS region
data "aws_region" "current" {}

# Get current AWS account ID
data "aws_caller_identity" "current" {}

# Example of an Redshift integration configuration
# Note: The source and target clusters must already exist
resource "awscc_redshift_integration" "example" {
  # ARN of the source Aurora cluster
  source_arn = "arn:aws:rds:${data.aws_region.current.name}:${data.aws_caller_identity.current.account_id}:cluster:example-aurora-cluster"

  # ARN of the target Redshift cluster
  target_arn = "arn:aws:redshift:${data.aws_region.current.name}:${data.aws_caller_identity.current.account_id}:cluster:example-redshift-cluster"

  # Name for the integration
  integration_name = "aurora-to-redshift"

  # Optional: KMS key for encryption (uncomment and provide key ID if needed)
  # kms_key_id = "arn:aws:kms:region:account:key/key-id"

  # Optional: Additional encryption context
  # additional_encryption_context = {
  #   "Environment" = "Production"
  # }

  tags = [{
    key   = "Modified By"
    value = "AWSCC"
  }]
}