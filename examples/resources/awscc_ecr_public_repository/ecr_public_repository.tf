resource "awscc_ecr_public_repository" "example" {
  repository_name = "example"
  tags = [{
    key   = "Managed By"
    value = "AWSCC"
  }]
}
