resource "awscc_ses_tenant" "example" {
  tenant_name = "example-tenant"
  
  tags = [{
    key   = "Environment"
    value = "test"
  }]
}
