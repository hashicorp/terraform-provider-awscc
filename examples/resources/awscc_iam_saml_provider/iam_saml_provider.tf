resource "awscc_iam_saml_provider" "example" {
  name                   = "example"
  saml_metadata_document = file("saml-metadata.xml")
  tags = [{
    key   = "Modified By"
    value = "AWSCC"
  }]

}