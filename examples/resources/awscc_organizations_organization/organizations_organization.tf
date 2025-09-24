# Create an AWS Organizations organization
resource "awscc_organizations_organization" "example" {
  # Feature set can be either "ALL" or "CONSOLIDATED_BILLING"
  feature_set = "ALL"
}
