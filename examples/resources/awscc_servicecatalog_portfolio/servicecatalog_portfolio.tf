resource "awscc_servicecatalog_portfolio" "example" {
  display_name  = "example-portfolio"
  provider_name = "Example Provider"
  description   = "Example Service Catalog Portfolio"

  tags = [
    {
      key   = "Environment"
      value = "example"
    },
    {
      key   = "Name"
      value = "example-portfolio"
    }
  ]
}
