resource "awscc_connect_contact_flow_module_version" "example" {
  contact_flow_module_arn = aws_connect_contact_flow_module.example.arn
  description             = "Example contact flow module version"
}
