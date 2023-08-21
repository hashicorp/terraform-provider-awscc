resource "awscc_ecr_public_repository" "example_catalog_data" {
  repository_name = "example-catalog-data"
  repository_catalog_data = {
    about_text             = "about text"
    architectures          = ["ARM"]
    operating_systems      = ["Linux"]
    repository_description = "Repository description"
    usage_text             = "Usage text"
  }
  tags = [{
    key   = "Managed By"
    value = "AWSCC"
  }]
}
